package json

import (
	"encoding/json"
	"hexagonal-go/shortener"

	"github.com/pkg/errors"
)

//Yes exactly, its so that we have a symbolic type with the interface on it that we can use to handle the redirects.
// That said, notice we take the empty Redirect struct in the json module and convert it to the redirect model from the shortner service.  By using an empty struct like that, we can generalize the data.  In other words, I could go into the model and change how the actual redirect struct looks like and we won't have to change the json or msgpack logic.
type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
