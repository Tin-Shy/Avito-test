package models

type TenderStatus string
type BidStatus string
type OrganizationType string

const (
    // Tender Status
    TenderCreated  TenderStatus = "CREATED"
    TenderPublished TenderStatus = "PUBLISHED"
    TenderClosed   TenderStatus = "CLOSED"

    // Bid Status
    BidCreated  BidStatus = "CREATED"
    BidPublished BidStatus = "PUBLISHED"
    BidCanceled BidStatus = "CANCELED"

    // Organization Type
    OrgIE  OrganizationType = "IE"
    OrgLLC OrganizationType = "LLC"
    OrgJSC OrganizationType = "JSC"
)
