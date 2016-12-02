package mongo

import (
    "github.com/dohrm/gotools/service"
    "gopkg.in/mgo.v2"
)

type Mongo struct {
    Url      string
    Database string
}

func NewMongo(url string, database string) *Mongo {
    return &Mongo{
        Url: url,
        Database: database,
    }
}

func MongoError(err error) *service.ServiceError {
    if err == mgo.ErrNotFound {
        return service.NewServiceError(err, "not.found", 404)
    } else if err != nil {
        return service.NewServiceError(err, "error.mongo", 500)
    }
    return nil
}

func (m *Mongo) Col(col string, fn func(*mgo.Collection) *service.ServiceError) *service.ServiceError {
    session, err := mgo.Dial(m.Url)
    if err != nil {
        return MongoError(err)
    }
    defer session.Close()
    c := session.DB(m.Database).C(col)
    return fn(c)
}

func (m *Mongo) Filter(col string, query interface{}, result interface{}, sorts ...string) service.ServiceFunction {
    return func() *service.ServiceError {
        return m.Col(col, func(c *mgo.Collection) *service.ServiceError {
            err := c.Find(query).Sort(sorts...).All(result)
            return MongoError(err)
        })
    }
}

func (m *Mongo) First(col string, query interface{}, result interface{}, sorts ...string) service.ServiceFunction {
    return func() *service.ServiceError {
        return m.Col(col, func(c *mgo.Collection) *service.ServiceError {
            err := c.Find(query).Sort(sorts...).One(result)
            return MongoError(err)
        })
    }
}

func (m *Mongo) Insert(col string, docs ...interface{}) service.ServiceFunction {
    return func() *service.ServiceError {
        return m.Col(col, func(c *mgo.Collection) *service.ServiceError {
            err := c.Insert(docs...)
            return MongoError(err)
        })
    }

}

func (m *Mongo) Update(col string, query interface{}, doc interface{}, result interface{}, sorts ...string) service.ServiceFunction {
    return func() *service.ServiceError {
        return m.Col(col, func(c *mgo.Collection) *service.ServiceError {
            err := c.Update(query, doc)
            if err != nil {
                return MongoError(err)
            }
            return m.First(col, query, result, sorts...)()
        })
    }
}