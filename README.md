# Go float64x4

A Go library that provides a vector type that mimics capabilities of common SIMD hardware
(with the hope that the Go compiler will optimize and actually translate this to SIMD instructions).

The creation of this Go library motivated by a (personal) long time desire to have first class vectors in programming languages,
and was inspired by the talk given here:
http://www.youtube.com/watch?v=CKh7UOELpPo

