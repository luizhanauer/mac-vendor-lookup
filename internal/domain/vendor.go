package domain

// VendorInfo represents the details of a specific MAC block (like IABs).
type VendorInfo struct {
	MacPrefix   string `json:"macPrefix"`
	CompanyName string `json:"companyName"`
	MaskLength  int    `json:"maskLength"`
}

// ShardAggregate groups the base /24 vendor and any specific sub-blocks under it.
type ShardAggregate struct {
	BaseVendor string       `json:"baseVendor,omitempty"`
	Blocks     []VendorInfo `json:"blocks,omitempty"`
}
