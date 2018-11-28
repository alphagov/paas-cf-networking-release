package migrations

var migration_v0058 = map[string][]string{
	"mysql": {
		`ALTER TABLE egress_policies ADD CONSTRAINT egress_policies_source_guid_destination_guid_app_lifecycle_unique UNIQUE (source_guid, destination_guid, app_lifecycle)`,
	},
	"postgres": {
		`ALTER TABLE egress_policies ADD CONSTRAINT egress_policies_source_guid_destination_guid_app_lifecycle_unique UNIQUE (source_guid, destination_guid, app_lifecycle)`,
	},
}
