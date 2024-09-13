package email

import (
	"fmt"
	"net"
	"strings"
)

func CheckDomain(domain string) error {

	if err := validateDomain(domain); err != nil {
		return err
	}

	if err := validateSPFRecords(domain); err != nil {
		return err
	}

	if err := validateDMARCRecords(domain); err != nil {
		return err
	}

	fmt.Println("All checks passed successfully\U0001F973")
	return nil
}

func validateDomain(domain string) error {
	fmt.Println("Checking domain...")

	// Check if domain exists
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return fmt.Errorf("domain does not exist")
	}

	fmt.Println("Domain exists \u2714")
	return nil
}

func validateSPFRecords(domain string) error {
	fmt.Println("Checking domain's SPF records...")

	// Check if domain has a TXT record
	txtRecords, err := net.LookupTXT(domain)
	if err != nil || len(txtRecords) == 0 {
		return fmt.Errorf("domain does not have a TXT record")
	}

	hasSPF := false
	for _, txtRecord := range txtRecords {
		if strings.HasPrefix(txtRecord, "v=spf1") {
			// Domain has a valid SPF record
			hasSPF = true
			break
		}
	}

	if !hasSPF {
		return fmt.Errorf("domain does not have a valid SPF record")
	} else {
		fmt.Println("Domain has a valid SPF record \u2714")
	}

	return nil
}

func validateDMARCRecords(domain string) error {
	fmt.Println("Checking domain's DMARC records...")

	// Check if domain has a DMARC record
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil || len(dmarcRecords) == 0 {
		return fmt.Errorf("domain does not have a DMARC record")
	}

	hasDMARC := false
	for _, dmarcRecord := range dmarcRecords {
		if strings.HasPrefix(dmarcRecord, "v=DMARC1") {
			// Domain has a valid DMARC record
			hasDMARC = true
			break
		}
	}

	if !hasDMARC {
		return fmt.Errorf("domain does not have a valid DMARC record")
	} else {
		fmt.Println("Domain has a valid DMARC record \u2714")
	}

	return nil
}
