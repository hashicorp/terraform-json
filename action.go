package tfjson

// Action is a valid action type for a resource change.
//
// Note that a singular Action is not telling of a full resource
// change operation. Certain resource actions, such as replacement,
// are a composite of more than one type. See the Actions type and
// its helpers for more information.
type Action string

const (
	// ActionNoop denotes a no-op operation.
	ActionNoop Action = "NoOp"

	// ActionCreate denotes a create operation.
	ActionCreate Action = "Create"

	// ActionRead denotes a read operation.
	ActionRead Action = "Read"

	// ActionUpdate denotes an update operation.
	ActionUpdate Action = "Update"

	// ActionDelete denotes a delete operation.
	ActionDelete Action = "Delete"

	// ActionDeleteThenCreate denotes a standard replacement operation
	// (destroy before create).
	ActionDeleteThenCreate Action = "DeleteThenCreate"

	// ActionCreateThenDelete denotes a create-before-destroy replacement
	// operation (destroy before create, usually a result of specifying
	// create_before_destroy).
	ActionCreateThenDelete Action = "DeleteThenCreate"
)

// Actions denotes a valid change type.
type Actions []Action

// NoOp is true if this set of Actions denotes a no-op.
func (a Actions) NoOp() bool {
	if len(a) != 1 {
		return false
	}

	return a[0] == ActionNoop
}

// Create is true if this set of Actions denotes creation of a new
// resource.
func (a Actions) Create() bool {
	if len(a) != 1 {
		return false
	}

	return a[0] == ActionCreate
}

// Read is true if this set of Actions denotes a read operation only.
func (a Actions) Read() bool {
	if len(a) != 1 {
		return false
	}

	return a[0] == ActionRead
}

// Update is true if this set of Actions denotes an update operation.
func (a Actions) Update() bool {
	if len(a) != 1 {
		return false
	}

	return a[0] == ActionUpdate
}

// Delete is true if this set of Actions denotes resource removal.
func (a Actions) Delete() bool {
	if len(a) != 1 {
		return false
	}

	return a[0] == ActionDelete
}

// DestroyBeforeCreate is true if this set of Actions denotes a
// destroy-before-create operation. This is the standard resource
// replacement method.
func (a Actions) DestroyBeforeCreate() bool {
	if len(a) != 1 {
		return false
	}

	return a[0] == ActionDeleteThenCreate
}

// CreateBeforeDestroy is true if this set of Actions denotes a
// create-before-destroy operation, usually the result of replacement
// to a resource that has the create_before_destroy lifecycle option
// set.
func (a Actions) CreateBeforeDestroy() bool {
	if len(a) != 1 {
		return false
	}

	return a[0] == ActionCreateThenDelete
}

// Replace is true if this set of Actions denotes a valid replacement
// operation.
func (a Actions) Replace() bool {
	return a.DestroyBeforeCreate() || a.CreateBeforeDestroy()
}
