//Nishant Arora
#include<bits/stdc++.h>
using namespace std;
unordered_map<int,int> m;
int main()
{
	int n,x,ans=0;
	cin>>n;
	for(int i=0;i<n;i++)
	{
		cin>>x;
		if(x!=0)
		{
			m[x]++;
			if(m[x]==1)ans++;
		}

	}
	cout<<ans;
}