package internal

/* wesley ebisuzaki v0.2
 *
 * takes 4 byte character string (single precision ieee big-endian)
 * and returns a float
 *
 * NaN, infinity are mapped into UNDEFINED.
 *
 * ansi C
 */
/*
 float32 ieee2flt_nan(unsigned char *ieee) {
	float64 fmant;
	int exp;

        if ((ieee[0] & 127) == 0 && ieee[1] == 0 && ieee[2] == 0 && ieee[3] == 0)
	   return (float) 0.0;

	exp = ((ieee[0] & 127) << 1) + (ieee[1] >> 7);

	if (exp == 255) return (float) UNDEFINED;

	fmant = (float64) ((int) ieee[3] + (int) (ieee[2] << 8) +
              (int) ((ieee[1] | 128) << 16));
	if (ieee[0] & 128) fmant = -fmant;


	return (float) (ldexp(fmant, (int) (exp - 128 - 22)));
}
*/

func ieee2flt_nan(ieee []byte) float32 {
	var fmant float64
	var exp int

	if (ieee[0]&127) == 0 && ieee[1] == 0 && ieee[2] == 0 && ieee[3] == 0 {
		return float32(0.0)
	}

	exp = (int(ieee[0]&127) << 1) + (int(ieee[1]) >> 7)

	if exp == 255 {
		return float32(UNDEFINED)
	}

	fmant = float64(int(ieee[3]) + (int(ieee[2]) << 8) +
		(int(ieee[1]|128) << 16))
	if (ieee[0] & 128) != 0 {
		fmant = -fmant
	}

	return float32(ldexp(fmant, int(exp-128-22)))
}
