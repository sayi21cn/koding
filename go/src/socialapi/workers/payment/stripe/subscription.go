package stripe

import (
	"socialapi/models/paymentmodel"

	"github.com/koding/bongo"
	stripe "github.com/stripe/stripe-go"
	stripeSub "github.com/stripe/stripe-go/sub"
)

func CreateSubscription(customer *paymentmodel.Customer, plan *paymentmodel.Plan) (*paymentmodel.Subscription, error) {
	subParams := &stripe.SubParams{
		Plan:     plan.ProviderPlanId,
		Customer: customer.ProviderCustomerId,
	}

	sub, err := stripeSub.New(subParams)
	if err != nil {
		return nil, err
	}

	subModel := paymentmodel.NewSubscription(sub.Id, ProviderName, plan, customer)
	err = subModel.Create()
	if err != nil {
		return nil, err
	}

	return subModel, nil
}

func FindCustomerActiveSubscriptions(customer *paymentmodel.Customer) ([]paymentmodel.Subscription, error) {
	var subs = []paymentmodel.Subscription{}

	if customer.Id == 0 {
		return nil, ErrCustomerIdIsNotSet
	}

	query := &bongo.Query{
		Selector: map[string]interface{}{
			"customer_id": customer.Id,
			"state":       "active",
		},
	}

	s := paymentmodel.Subscription{}
	err := s.Some(&subs, query)
	if err != nil {
		return nil, err
	}

	return subs, nil
}

func CancelSubscription(customer *paymentmodel.Customer, subscription paymentmodel.Subscription) error {
	return nil
}
