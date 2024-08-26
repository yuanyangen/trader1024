package mongo

import (
	"context"
	"dario.cat/mergo"
	"errors"
	"fmt"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionNameSubjects = "subject_infos"

func SaveSubjectInfo(ctx context.Context, m *model.SubjectDO) error {
	if m == nil {
		return fmt.Errorf("m nil")
	}
	if m.CNName == "" {
		return fmt.Errorf("cn name empty")
	}
	subject, err := QuerySubject(ctx, m.CNName)
	if err != nil {
		return err
	}
	if subject == nil {
		return InsertSubjectInfo(ctx, m)
	} else {
		mergo.Merge(m, subject, mergo.WithOverride)

		return UpdateSubject(ctx, m.CNName, m)
	}
}

func UpdateSubject(ctx context.Context, cnName string, s *model.SubjectDO) error {
	collection := Trader1024Db.Collection(collectionNameSubjects)
	_, err := collection.UpdateOne(ctx, bson.D{{"cnname", cnName}}, bson.D{{"$set", s}})
	if err != nil {
		logs.Info("insert error %v %v", err, cnName)
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func InsertSubjectInfo(ctx context.Context, m *model.SubjectDO) error {
	collection := Trader1024Db.Collection(collectionNameSubjects)
	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		logs.Info("insert error %v %v\n", err, m.String())
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
func QuerySubject(ctx context.Context, cnName string) (*model.SubjectDO, error) {
	collection := Trader1024Db.Collection(collectionNameSubjects)
	var result *model.SubjectDO
	err := collection.FindOne(ctx, bson.D{{"cnname", cnName}}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logs.Info("No document was found with the cnName  %s\n", cnName)
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func QueryAllSubject(ctx context.Context) ([]*model.SubjectDO, error) {
	collection := Trader1024Db.Collection(collectionNameSubjects)
	var result []*model.SubjectDO
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logs.Info("No document was found \n")
			return nil, nil
		}
		return nil, err
	}
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
