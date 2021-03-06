package waiter

import (
	"time"

	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	// Maximum amount of time to wait for an AdminAccount to return Enabled
	AdminAccountEnabledTimeout = 5 * time.Minute

	// Maximum amount of time to wait for an AdminAccount to return NotFound
	AdminAccountNotFoundTimeout = 5 * time.Minute

	// Maximum amount of time to wait for Standards Subscriptions to return READY
	StandardsSubscriptionsReadyTimeout = 5 * time.Minute

	// Maximum amount of time to wait for Standards Subscriptions to be Deleted
	StandardsSubscriptionsDeletedTimeout = 5 * time.Minute
)

// AdminAccountEnabled waits for an AdminAccount to return Enabled
func AdminAccountEnabled(conn *securityhub.SecurityHub, adminAccountID string) (*securityhub.AdminAccount, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{AdminStatusNotFound},
		Target:  []string{securityhub.AdminStatusEnabled},
		Refresh: AdminAccountAdminStatus(conn, adminAccountID),
		Timeout: AdminAccountEnabledTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*securityhub.AdminAccount); ok {
		return output, err
	}

	return nil, err
}

// AdminAccountNotFound waits for an AdminAccount to return NotFound
func AdminAccountNotFound(conn *securityhub.SecurityHub, adminAccountID string) (*securityhub.AdminAccount, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{securityhub.AdminStatusDisableInProgress},
		Target:  []string{AdminStatusNotFound},
		Refresh: AdminAccountAdminStatus(conn, adminAccountID),
		Timeout: AdminAccountNotFoundTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*securityhub.AdminAccount); ok {
		return output, err
	}

	return nil, err
}

func StandardsSubscriptionsReady(conn *securityhub.SecurityHub) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{securityhub.StandardsStatusIncomplete, securityhub.StandardsStatusPending, StandardsStatusNotFound},
		Target:  []string{securityhub.StandardsStatusReady},
		Refresh: StandardsSubscriptionsStatus(conn),
		Timeout: StandardsSubscriptionsReadyTimeout,
	}

	_, err := stateConf.WaitForState()

	return err
}

func StandardsSubscriptionsDeleted(conn *securityhub.SecurityHub) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{securityhub.StandardsStatusDeleting},
		Target:  []string{StandardsStatusNotFound},
		Refresh: StandardsSubscriptionsStatus(conn),
		Timeout: StandardsSubscriptionsDeletedTimeout,
	}

	_, err := stateConf.WaitForState()

	return err
}
