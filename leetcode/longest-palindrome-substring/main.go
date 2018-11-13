package main

func longestPalindrome(s string) (maxp string) {
	n := len(s)
	i := 0
	for i < n {
		il, iu := i-1, i
		localp := s[i : i+1]
		for iu < n && s[i] == s[iu] {
			// fmt.Printf("1 -- s=%q,i=%d,il=%d,iu=%d,maxp=%q,localp=%q\n", s, i, il, iu, maxp, localp)
			localp = s[i : iu+1]
			iu++
		}

		// fmt.Printf("2 -- s=%q,i=%d,il=%d,iu=%d,maxp=%q,localp=%q\n", s, i, il, iu, maxp, localp)

		for il >= 0 && iu < n && s[il] == s[iu] {
			// fmt.Printf("3 -- s=%q,i=%d,il=%d,iu=%d,maxp=%q,localp=%q\n", s, i, il, iu, maxp, localp)
			localp = s[il : iu+1]
			il--
			iu++
		}

		if len(localp) > len(maxp) {
			maxp = localp
		}

		// fmt.Printf("4 -- s=%q,i=%d,il=%d,iu=%d,maxp=%q,localp=%q\n", s, i, il, iu, maxp, localp)

		i++
	}

	return
}
