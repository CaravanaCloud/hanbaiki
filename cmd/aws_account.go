package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/charmbracelet/log"
)

// AWSAccount implements the ResourceOperations interface for AWS accounts
type AWSAccount struct{}

// Allocate implements ResourceType.
func (a AWSAccount) Allocate() error {
	fmt.Println("Allocate AWS Account...")
	return nil

}

// Claim implements ResourceType.
func (a AWSAccount) Claim() error {
	fmt.Println("Claim AWS Account...")
	return nil

}

// Cleanup implements ResourceType.
func (a AWSAccount) Cleanup() error {
	fmt.Println("Cleanup AWS Account...")
	return nil

}

// Dispose implements ResourceType.
func (a AWSAccount) Dispose() error {
	fmt.Println("Dispose AWS Account...")
	return nil

}

// Lease implements ResourceType.
func (a AWSAccount) Lease() error {
	fmt.Println("Lease AWS Account...")
	return nil

}

// findOUByName finds the Organizational Unit ID by name
func findOUByName(svc *organizations.Organizations, ouName string) (string, error) {
	input := &organizations.ListRootsInput{}

	// Assume single root scenario
	root, err := svc.ListRoots(input)
	if err != nil {
		return "", err
	}
	rootID := root.Roots[0].Id // Consider the first root

	// List OUs for the root
	ouInput := &organizations.ListOrganizationalUnitsForParentInput{
		ParentId: rootID,
	}
	var foundOUId string // Variable to store the found OU ID
	err = svc.ListOrganizationalUnitsForParentPages(ouInput, func(page *organizations.ListOrganizationalUnitsForParentOutput, lastPage bool) bool {
		for _, ou := range page.OrganizationalUnits {
			if *ou.Name == ouName {
				foundOUId = *ou.Id // Store OU ID when found
				return false       // Stop paging
			}
		}
		return !lastPage // Continue paging if not found
	})

	if err != nil {
		return "", err
	}

	if foundOUId == "" {
		return "", fmt.Errorf("OU with name %s not found", ouName)
	}

	return foundOUId, nil // Return the found OU ID
}

// listAccountsForOU lists AWS accounts for a specific OU
func listAccountsForOU(svc *organizations.Organizations, ouID string) error {
	var accountsInfo []string
	err := svc.ListAccountsForParentPages(&organizations.ListAccountsForParentInput{ParentId: &ouID},
		func(page *organizations.ListAccountsForParentOutput, lastPage bool) bool {
			for _, account := range page.Accounts {
				info := fmt.Sprintf("Account ID: %s, Name: %s, Email: %s", *account.Id, *account.Name, *account.Email)
				accountsInfo = append(accountsInfo, info)
			}
			return !lastPage
		},
	)
	if err != nil {
		log.Fatalf("Failed to list accounts in OU: %s", err)
	}

	for _, accountInfo := range accountsInfo {
		fmt.Println(accountInfo)
	}
	return nil
}

// listAllAccounts lists all AWS accounts
func listAllAccounts(svc *organizations.Organizations) error {
	var accountsInfo []string
	err := svc.ListAccountsPages(&organizations.ListAccountsInput{},
		func(page *organizations.ListAccountsOutput, lastPage bool) bool {
			for _, account := range page.Accounts {
				info := fmt.Sprintf("Account ID: %s, Name: %s, Email: %s", *account.Id, *account.Name, *account.Email)
				accountsInfo = append(accountsInfo, info)
			}
			return !lastPage
		},
	)
	if err != nil {
		log.Fatalf("Failed to list accounts: %s", err)
	}

	for _, accountInfo := range accountsInfo {
		fmt.Println(accountInfo)
	}
	return nil
}

// List implements ResourceType.
func (a AWSAccount) List(data map[string]string) error {
	ou := data["ou"]
	log.Infof("List AWS Accounts for OU[%s].", ou)
	sess, err := session.NewSession(&aws.Config{})

	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}

	svc := organizations.New(sess)
	if ou != "" {
		// Find OU ID by name
		ouID, err := findOUByName(svc, ou)
		if err != nil {
			log.Warnf("Failed to find OU by name[%s]: %v", ou, err)
			return err
		}

		// List accounts in the found OU
		return listAccountsForOU(svc, ouID)
	}

	// Otherwise, list all accounts
	return listAllAccounts(svc)

}

// Prepare implements ResourceType.
func (a AWSAccount) Prepare() error {
	fmt.Println("Prepare AWS Account...")
	return nil
}
