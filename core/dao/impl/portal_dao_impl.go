/**
 * This file is auto generated by tto v0.4.5 !
 * If you want to modify this code, please read the guide
 * to modify code template.
 *
 * Get started: https://github.com/ixre/tto
 *
 * Copyright (C) 2021 56X.NET, All rights reserved.
 *
 * name : portal_nav_dao_impl.go
 * author : jarrysix
 * date : 2021/12/12 09:09:36
 * description :
 * history :
 */
package impl

import (
	"database/sql"
	"fmt"
	"github.com/ixre/gof/db"
	"github.com/ixre/gof/db/orm"
	"go2o/core/dao"
	"go2o/core/dao/model"
	"log"
)

var _ dao.IPortalDao = new(portalNavDaoImpl)

type portalNavDaoImpl struct {
	_orm orm.Orm
}

var portalNavDaoImplMapped = false

// Create new PortalNavDao
func NewPortalDao(o orm.Orm) dao.IPortalDao {
	if !portalNavDaoImplMapped {
		_ = o.Mapping(model.PortalNav{}, "portal_nav")
		portalNavDaoImplMapped = true
	}
	return &portalNavDaoImpl{
		_orm: o,
	}
}

// Get 门户导航
func (p *portalNavDaoImpl) GetNav(primary interface{}) *model.PortalNav {
	e := model.PortalNav{}
	err := p._orm.Get(primary, &e)
	if err == nil {
		return &e
	}
	if err != sql.ErrNoRows {
		log.Println("[ Orm][ Error]:", err.Error(), "; Entity:PortalNav")
	}
	return nil
}

// GetBy 门户导航
func (p *portalNavDaoImpl) GetNavBy(where string, v ...interface{}) *model.PortalNav {
	e := model.PortalNav{}
	err := p._orm.GetBy(&e, where, v...)
	if err == nil {
		return &e
	}
	if err != sql.ErrNoRows {
		log.Println("[ Orm][ Error]:", err.Error(), "; Entity:PortalNav")
	}
	return nil
}

// Count 门户导航 by condition
func (p *portalNavDaoImpl) CountNav(where string, v ...interface{}) (int, error) {
	return p._orm.Count(model.PortalNav{}, where, v...)
}

// Select 门户导航
func (p *portalNavDaoImpl) SelectNav(where string, v ...interface{}) []*model.PortalNav {
	list := make([]*model.PortalNav, 0)
	err := p._orm.Select(&list, where, v...)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[ Orm][ Error]:", err.Error(), "; Entity:PortalNav")
	}
	return list
}

// Save 门户导航
func (p *portalNavDaoImpl) SaveNav(v *model.PortalNav) (int, error) {
	id, err := orm.Save(p._orm, v, int(v.Id))
	if err != nil && err != sql.ErrNoRows {
		log.Println("[ Orm][ Error]:", err.Error(), "; Entity:PortalNav")
	}
	return id, err
}

// Delete 门户导航
func (p *portalNavDaoImpl) DeleteNav(primary interface{}) error {
	err := p._orm.DeleteByPk(model.PortalNav{}, primary)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[ Orm][ Error]:", err.Error(), "; Entity:PortalNav")
	}
	return err
}

// Batch Delete 门户导航
func (p *portalNavDaoImpl) BatchDeleteNav(where string, v ...interface{}) (int64, error) {
	r, err := p._orm.Delete(model.PortalNav{}, where, v...)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[ Orm][ Error]:", err.Error(), "; Entity:PortalNav")
	}
	return r, err
}

// Query paging data
func (p *portalNavDaoImpl) PagingQueryNav(begin, end int, where, orderBy string) (total int, rows []map[string]interface{}) {
	if orderBy != "" {
		orderBy = "ORDER BY " + orderBy
	}
	if where == "" {
		where = "1=1"
	}
	s := fmt.Sprintf(`SELECT COUNT(0) FROM portal_nav WHERE %s`, where)
	_ = p._orm.Connector().ExecScalar(s, &total)
	if total > 0 {
		s = fmt.Sprintf(`SELECT * FROM portal_nav WHERE %s %s
	        LIMIT $2 OFFSET $1`,
			where, orderBy)
		err := p._orm.Connector().Query(s, func(_rows *sql.Rows) {
			rows = db.RowsToMarshalMap(_rows)
		}, begin, end-begin)
		if err != nil {
			log.Println(fmt.Sprintf("[ Orm][ Error]: %s (table:portal_nav)", err.Error()))
		}
	} else {
		rows = make([]map[string]interface{}, 0)
	}
	return total, rows
}
