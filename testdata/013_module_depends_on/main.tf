# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

module "foo" {
    source = "vancluever/module/null"

    depends_on = [
        null_resource.bar
    ]
}

resource "null_resource" "bar" {}