package main

import (
	"context"
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
)

// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	var versionSlice []*semver.Version
	// This is just an example structure of the code, if you implement this interface, the test cases in main_test.go are very easy to run

	//simple way

	/* 
	Error if there is something like{
		"2.8.2", "2.1.6"
	}
	it will not output [2.8.2]
	it will output [2.8.2, 2.1.6]
	*/

	if(len(releases) == 0){
		return nil
	}else  if(len(releases)==1){
		if(releases[0].LessThan(*minVersion)){
			return nil
		}else{
			return releases
		}
	}
	for _, r:= range releases{
		if(!r.LessThan(*minVersion)){
			versionSlice = append(versionSlice, r)
		}
	}
	semver.Sort(versionSlice)

	var temp []*semver.Version
	greatestVer := versionSlice[len(versionSlice)-1]
	lower := LowerBound(greatestVer)
	temp = append(temp, greatestVer)

	for i := len(versionSlice)-1; i >= 0; i--{
		if( !lower.LessThan(*versionSlice[i])&&*lower!=*versionSlice[i]){
			greatestVer := versionSlice[i]
			lower = LowerBound(greatestVer)
			temp = append(temp, greatestVer)
		}else{
		}
		// the print functions below are for debug
		// fmt.Printf("%d  %s\n",i, temp)
	}
	// fmt.Printf("\n")
	return temp
}


func LowerBound(version *semver.Version) *semver.Version{
	temp := version.String()

	for i := len(temp)-1; i >= 0; i-- {
		if(temp[i] == '.'){
			temp = string(temp[:i+1]+"0")
			break
		}
	}
	return semver.New(temp)
}

// Here we implement the basics of communicating with github through the library as well as printing the version
// You will need to implement LatestVersions function as well as make this application support the file format outlined in the README
// Please use the format defined by the fmt.Printf line at the bottom, as we will define a passing coding challenge as one that outputs
// the correct information, including this line
func main() {
	// Github
	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.ListOptions{PerPage: 10}
	releases, _, err := client.Repositories.ListReleases(ctx, "kubernetes", "kubernetes", opt)
	if err != nil {
		panic(err) // is this really a good way?
	}
	minVersion := semver.New("1.8.0")
	allReleases := make([]*semver.Version, len(releases))
	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}
		allReleases[i] = semver.New(versionString)
	}
	versionSlice := LatestVersions(allReleases, minVersion)

	fmt.Printf("latest versions of kubernetes/kubernetes: %s", releases)

	//my implementation
	opt = &github.ListOptions{PerPage: 10}
	releases, _, err = client.Repositories.ListReleases(ctx, "prometheus", "prometheus", opt)
	if err != nil {
		panic(err) // is this really a good way?
		//not a good way
		//this function should retrun err to notify the caller function 
		//but in here we use main, so no caller
		//I do not implement the err notification.
	}
	minVersion = semver.New("2.2.0")
	allReleases = make([]*semver.Version, len(releases))
	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}
		allReleases[i] = semver.New(versionString)
	}
	versionSlice = LatestVersions(allReleases, minVersion)
	fmt.Printf("\nlatest versions of prometheus/prometheus: %s", versionSlice)

}



/*
{
   "github_url": "https://gist.github.com/Peterllau/10d7b23adddd910ddb88e50ea25b5191",
   "contact_email": "cwipac@connect.ust.hk"
}
*/