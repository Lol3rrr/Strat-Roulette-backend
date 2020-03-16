package database

// SessionInterface provides a simple interface to do all the needed
// database related actions
type SessionInterface interface {
	// Get simply loads the first entry matching the query params and
	// loads it into the result interface param
	Get(map[string]interface{}, interface{}) error
	// GetAll loads all the entrys matching the query params
	GetAll(map[string]interface{}, interface{}) error
	// Insert simply inserts the given interface into the Database
	Insert(interface{}) error
	// Update takes the given filter and updateing params and uses them to
	// update all the matching entrys
	Update(map[string]interface{}, map[string]interface{}) error
	// Delete takes the given filter params and deletes the first entry matching it
	Delete(map[string]interface{}) error
	// DeleteMany takes the given filter params and deletes every entry matching it
	DeleteMany(map[string]interface{}) error
}
