package links

import (
	"database/sql"
	syserrors "link_shortener/errors"

	"github.com/mattn/go-sqlite3"
)

const duplicatePrimaryKeyErrorCode = 1555
const uniqueConstraintErrorCode = 2067

type SqliteLinkDA struct {
	database     *sql.DB
	hashFunction Hasher
	errorHandler syserrors.ErrorHandler
}

func CreateSQLiteLinkDa(db *sql.DB, hashFunction Hasher, errorHandler syserrors.ErrorHandler) *SqliteLinkDA {
	linkDa := new(SqliteLinkDA)
	linkDa.database = db
	linkDa.hashFunction = hashFunction
	linkDa.errorHandler = errorHandler

	return linkDa
}

func (da *SqliteLinkDA) Get(id string) (*Link, error) {
	return da.takeOneQueryResult("SELECT id, external_url FROM links WHERE id=? LIMIT 1", id)
}

func (da *SqliteLinkDA) GetByUrl(url string) (*Link, error) {
	return da.takeOneQueryResult("SELECT id, external_url FROM links WHERE external_url=? LIMIT 1", url)
}

func (da *SqliteLinkDA) takeOneQueryResult(query string, args ...any) (*Link, error) {
	var returnLink *Link = nil

	var rows *sql.Rows
	var err error

	//Prepare the query
	statement, err := da.database.Prepare(query)
	if err != nil {
		return nil, da.sqlErrorResponse(err)
	}
	defer statement.Close()

	//Execute the prepared statement
	rows, err = statement.Query(args...)
	if err != nil {
		return nil, da.sqlErrorResponse(err)
	}
	defer rows.Close()

	//Load the row if there is any
	if !rows.Next() {
		return nil, ErrNotFound{}
	}

	//Parse link from the query
	returnLink = new(Link)
	err = rows.Scan(&returnLink.Id, &returnLink.ExternalURL)
	if err != nil {
		return nil, da.sqlErrorResponse(err)
	}

	return returnLink, nil
}

func (da *SqliteLinkDA) Save(link Link) (*Link, error) {
	return da.saveRecursive(link, 0, 10)
}

func (da *SqliteLinkDA) saveRecursive(link Link, iteration int, maxRecursionDepth int) (*Link, error) {
	if iteration >= maxRecursionDepth {
		return nil, ErrUniqueIdGenerationFailed{}
	}

	var returnLink *Link = nil
	var returnError error = nil

	sqlError := da.sqlExec("INSERT INTO links(id, external_url) VALUES(?, ?)", link.Id, link.ExternalURL)

	if sqlError == nil {
		returnLink = new(Link)
		*returnLink = link

	} else if isDuplicateIdError(sqlError) {
		link.Id = da.hashFunction.Hash(link.Id)
		returnLink, returnError = da.saveRecursive(link, iteration+1, maxRecursionDepth)

	} else if isDuplicateUrlError(sqlError) {
		returnError = ErrUrlDuplicate{}
	} else {
		returnError = da.sqlErrorResponse(sqlError)
	}

	return returnLink, returnError
}

func (da *SqliteLinkDA) sqlExec(query string, values ...any) error {
	//Prepare the query
	statement, err := da.database.Prepare(query)
	if err != nil {
		return da.sqlErrorResponse(err)
	}
	defer statement.Close()

	//Execute the prepared statement
	_, err = statement.Exec(values...)

	return err
}

func (da *SqliteLinkDA) sqlErrorResponse(err error) error {
	da.errorHandler.InternalDAError(err)
	return ErrRepositoryError{}
}

func isDuplicateIdError(err error) bool {
	return err.(sqlite3.Error).ExtendedCode == duplicatePrimaryKeyErrorCode
}

func isDuplicateUrlError(err error) bool {
	return err.(sqlite3.Error).ExtendedCode == uniqueConstraintErrorCode
}
