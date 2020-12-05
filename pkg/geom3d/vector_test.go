package geom3d

import "testing"

func TestVector_Dot(t *testing.T) {
	type fields struct {
		X Fl
		Y Fl
		Z Fl
	}
	type args struct {
		b Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Fl
	}{
		{
			name:   "when the vectors are orhogonal, dot should be 0",
			fields: fields{1,0,0},
			args:   args{
				b: Vector{0,0,0},
			},
			want:   0,
		}, {
			name:   "when vectors are opposite, doe sohuld be -1",
			fields: fields{-1,0, 0},
			args:   args{
				b: Vector{1,0,0},
			},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				a := Vector{
					X: tt.fields.X,
					Y: tt.fields.Y,
					Z: tt.fields.Z,
				}
				if got := a.Dot(tt.args.b); got != tt.want {
					t.Errorf("Dot() = %v, want %v", got, tt.want)
				}
			})
	}
}
