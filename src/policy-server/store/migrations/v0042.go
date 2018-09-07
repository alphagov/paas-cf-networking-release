package migrations

var migration_v0042 = map[string][]string{
	"mysql": {
		`ALTER TABLE egress_policies
		 ADD CONSTRAINT egress_policies_destination_guid_fk FOREIGN KEY (destination_guid)
		 REFERENCES terminals(guid),
		 MODIFY destination_guid VARCHAR(36) NOT NULL UNIQUE,
		 DROP FOREIGN KEY egress_policies_destination_id_fk,
		 DROP destination_id;`,
	},
	"postgres": {
		`ALTER TABLE egress_policies
		 ADD CONSTRAINT egress_policies_destination_guid_fk FOREIGN KEY (destination_guid) REFERENCES terminals(guid),
		 ADD CONSTRAINT egress_policies_destination_guid_unique UNIQUE (destination_guid),
		 ALTER COLUMN destination_guid SET NOT NULL,
		 DROP destination_id;`,
	},
}
