package mongodb

import (
	"context"
	"movies-service/internal/core/domain/entity"
	"movies-service/internal/core/port/output"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Movie struct {
	collection *mongo.Collection
}

type movieRepository struct {
	Id    int    `bson:"_id"`
	Title string `bson:"title"`
	Year  string `bson:"year"`
}

func NewMovieRepository(collection *mongo.Collection) *Movie {
	return &Movie{
		collection: collection,
	}
}

func ToDocument(movie *entity.MovieEntity) movieRepository {
	return movieRepository{
		Id:    movie.GetID(),
		Title: movie.GetTitle(),
		Year:  movie.GetYear(),
	}
}

func ToDomain(doc movieRepository) (*entity.MovieEntity, error) {
	return entity.NewMovieEntity(doc.Id, doc.Title, doc.Year)
}

func (m *Movie) GetMovieByID(ctx context.Context, id int) (*entity.MovieEntity, error) {
	filter := bson.M{"_id": id}

	var doc movieRepository
	err := m.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		return nil, err
	}

	return ToDomain(doc)
}

func (m *Movie) ListMovies(ctx context.Context, filters output.Listfilters, pagination output.Pagination, sorting output.Sorting) ([]entity.MovieEntity, error) {
	filter := bson.M{}
	if filters.Title != "" {
		filter["title"] = bson.M{"$regex": filters.Title, "$options": "i"}
	}
	if filters.Year != "" {
		filter["year"] = filters.Year
	}

	findOptions := options.Find()
	if pagination.Limit > 0 {
		findOptions.SetLimit(int64(pagination.Limit))
	}
	if pagination.Page > 0 && pagination.Limit > 0 {
		findOptions.SetSkip(int64((pagination.Page - 1) * pagination.Limit))
	}
	if sorting.SortBy != "" {
		findOptions.SetSort(bson.D{bson.E{Key: sorting.SortBy, Value: 1}})
	}

	var docs []movieRepository
	cursor, err := m.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc movieRepository
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	var movies []entity.MovieEntity
	for _, doc := range docs {
		movie, err := ToDomain(doc)
		if err != nil {
			return nil, err
		}
		movies = append(movies, *movie)
	}

	return movies, nil
}

func (m *Movie) CountMovies(ctx context.Context, filters output.Listfilters) (int, error) {
	filter := bson.M{}
	if filters.Title != "" {
		filter["title"] = bson.M{"$regex": filters.Title, "$options": "i"}
	}
	if filters.Year != "" {
		filter["year"] = filters.Year
	}

	count, err := m.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (m *Movie) CreateMovie(ctx context.Context, movie *entity.MovieEntity) (*entity.MovieEntity, error) {
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	filter := bson.M{"_id": movie.GetID()}
	update := bson.M{"$set": ToDocument(movie)}

	var updatedDoc movieRepository
	err := m.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDoc)
	if err != nil {
		return nil, err
	}

	return ToDomain(updatedDoc)
}

func (m *Movie) DeleteMovie(ctx context.Context, id int) error {
	filter := bson.M{"_id": id}

	result, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
