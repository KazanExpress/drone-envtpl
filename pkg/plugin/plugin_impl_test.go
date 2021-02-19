// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"os"
	"reflect"
	"testing"
)

func Test_fillJinjaTemplate(t *testing.T) {
	type args struct {
		template []byte
	}
	os.Setenv("TOP", "chik")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"simple test a->a", args{[]byte("kek")}, []byte("kek"), false},
		{"test with jinja var", args{[]byte("{%- set k = 'kek' -%}{{k}}top")}, []byte("kektop"), false},
		{"test with env vars", args{[]byte("{{TOP}}")}, []byte("chik"), false},
		{"should fail", args{[]byte("{{kek}")}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fillJinjaTemplate(tt.args.template)
			if (err != nil) != tt.wantErr {
				t.Errorf("fillTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
