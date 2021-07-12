package stack

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	stackOK := "stack ok"

	cases := []struct {
		name         string
		err          error
		expErr       error
		expErrString string
	}{
		{
			name:   "err is nil",
			err:    nil,
			expErr: nil,
		},
		{
			name: "err already contains stack",
			err: fmt.Errorf(
				"something went wrong: %w",
				&StackError{
					Stack: "abc",
					Err:   errors.New("inner error"),
				},
			),
			expErr: fmt.Errorf(
				"something went wrong: %w",
				&StackError{
					Stack: "abc",
					Err:   errors.New("inner error"),
				},
			),
			expErrString: "something went wrong: inner error\nabc",
		},
		{
			name: "ok",
			err:  fmt.Errorf("something went wrong: %w", errors.New("inner error")),
			expErr: &StackError{
				Err:   fmt.Errorf("something went wrong: %w", errors.New("inner error")),
				Stack: stackOK,
			},
			expErrString: "something went wrong: inner error\nstack ok",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			getStackFunc := func() []byte {
				return []byte(stackOK)
			}

			err := wrap(getStackFunc, tc.err)
			require.Equal(t, tc.expErr, err)
			if err != nil {
				require.Equal(t, tc.expErrString, err.Error())
			}
		})
	}
}
