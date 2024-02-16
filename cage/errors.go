package cage

import (
	"fmt"
)

var (
	// ErrBreak @todo doc
	ErrBreak = fmt.Errorf("break propagation")

	// ErrNilPointer @todo doc
	ErrNilPointer = fmt.Errorf("nil pointer")

	// ErrDuplicateChild @todo doc
	ErrDuplicateChild = fmt.Errorf("duplicate child")

	// ErrChildNotFound @todo doc
	ErrChildNotFound = fmt.Errorf("child not found")

	// ErrDuplicateScene @todo doc
	ErrDuplicateScene = fmt.Errorf("duplicate scene")

	// ErrSceneNotFound @todo doc
	ErrSceneNotFound = fmt.Errorf("scene not found")

	// ErrSceneQueued @todo doc
	ErrSceneQueued = fmt.Errorf("scene is queued to run")

	// ErrSceneRunning @todo doc
	ErrSceneRunning = fmt.Errorf("scene in running")

	// ErrResAlreadyQueued @todo doc
	ErrResAlreadyQueued = fmt.Errorf("resource already queued")

	// ErrImageNotFound @todo doc
	ErrImageNotFound = fmt.Errorf("image not found")
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %s", ErrNilPointer, arg)
}

func errDuplicateChild(node, child string) error {
	return fmt.Errorf("%w : %s->%s", ErrDuplicateChild, node, child)
}

func errChildNotFound(node, child string) error {
	return fmt.Errorf("%w : %s->%s", ErrChildNotFound, node, child)
}

func errDuplicateScene(id string) error {
	return fmt.Errorf("%w : %s", ErrDuplicateScene, id)
}

func errSceneNotFound(id string) error {
	return fmt.Errorf("%w : %s", ErrSceneNotFound, id)
}

func errSceneQueued(id string) error {
	return fmt.Errorf("%w : %s", ErrSceneQueued, id)
}

func errSceneRunning(id string) error {
	return fmt.Errorf("%w : %s", ErrSceneRunning, id)
}

func errResAlreadyQueued(id string) error {
	return fmt.Errorf("%w : %s", ErrResAlreadyQueued, id)
}

func errImageNotFound(id string) error {
	return fmt.Errorf("%w : %s", ErrImageNotFound, id)
}
