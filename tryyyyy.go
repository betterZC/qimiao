package main


func qm(s string) string {
	res := ""
	dp := make([][]bool, len(s))
	for i,_ :=range s {
		dp[i] = make([]bool, len(s))
	}
	for i:=len(s)-1;i>=0;i--{
		for j:=i;j<len(s);j++ {
			//dp[i][j] = (s[i]==s[j]) && (dp[i+1][j-1] || j-i<3)//???????
			dp[i][j] = (s[i]==s[j]) && (j-i<3 || dp[i+1][j-1])
			//dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res=="" || ((j-i+1)>len(res)) ) {
				res = s[i:j+1]
			}
		}
	}
	return res
}

func main() {
	//s := "abc"
	//for i,j :=range(s) {
	//	println(i)
	//	print(string(j))
	//}
	//println("------------")
	//for i := 0; i < len(s); i++ {
	//	println(i)
	//}
	qm("abcbmghma")
}