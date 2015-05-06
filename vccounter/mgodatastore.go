package vccounter

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	AppVersionCodeCollectionName = "versionCodes"
	AppIdKey                     = "appid"
)

type AppVersionCodeMGODataStore struct {
	mgoSession *mgo.Session
}

func (avcDataStore *AppVersionCodeMGODataStore) OpenConnection(url string) error {

	avcDataStore.CloseConnection()

	session, err := mgo.Dial(url)
	if err != nil {

		return errors.New("Problem connecting to database")
	}

	avcDataStore.mgoSession = session

	return nil
}

func (avcDataStore *AppVersionCodeMGODataStore) CloseConnection() {

	if avcDataStore.mgoSession != nil {

		avcDataStore.mgoSession.Close()
		avcDataStore = nil
	}
}

func (avcDataStore *AppVersionCodeMGODataStore) initDBSession() (*mgo.Session, *mgo.Collection) {

	session := avcDataStore.mgoSession.Copy()
	c := session.DB("").C(AppVersionCodeCollectionName)

	return session, c
}

func (avcDataStore *AppVersionCodeMGODataStore) DeleteApp(appId string) error {

	session, c := avcDataStore.initDBSession()
	defer session.Close()

	err := c.Remove(bson.M{AppIdKey: appId})

	if err != nil {
		return err
	}

	return nil
}

func (avcDataStore *AppVersionCodeMGODataStore) UpdateAppVersionCode(ac *AppCode) error {

	session, c := avcDataStore.initDBSession()
	defer session.Close()

	_, err := c.Upsert(bson.M{AppIdKey: ac.AppId}, ac)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (avcDataStore *AppVersionCodeMGODataStore) CurrentAppVersionCode(appId string) (*AppCode, error) {

	ac := AppCode{appId, 0}

	session, c := avcDataStore.initDBSession()
	defer session.Close()

	query := c.Find(bson.M{AppIdKey: appId})

	nEntries, err := query.Count()

	if err != nil {
		return nil, err
	}

	if nEntries == 0 {
		err = c.Insert(&ac)
	} else {
		err = query.One(&ac)
	}

	if err != nil {
		return nil, err
	}

	return &ac, nil
}
