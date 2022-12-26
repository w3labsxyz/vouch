// Copyright © 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock

import (
	"context"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/attestantio/vouch/services/beaconblockproposer"
	e2wtypes "github.com/wealdtech/go-eth2-wallet-types/v2"
)

// Service is a mock block relay.
type Service struct{}

// New creates a new mock block relay.
func New() *Service {
	return &Service{}
}

// SubmitValidatorRegistrations submits validator registrations for the given accounts.
func (*Service) SubmitValidatorRegistrations(_ context.Context, _ map[phase0.ValidatorIndex]e2wtypes.Account) error {
	return nil
}

// ProposerConfig returns the proposer configuration for the given validator.
func (*Service) ProposerConfig(_ context.Context,
	_ e2wtypes.Account,
	_ phase0.BLSPubKey,
) (
	*beaconblockproposer.ProposerConfig,
	error,
) {
	return nil, nil
}
