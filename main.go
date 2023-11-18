package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeRequest(roomID string) {
	url := "https://webcast16-normal-c-alisg.tiktokv.com/webcast/room/like/?enter_from=live_merge-live_cover"

	requestBody := []byte(fmt.Sprintf("room_id=%s&count=15&enter_from_merge=live_merge&enter_method=live_cover&is_ad=0&label=user_server_live_like&tag=live_ad", roomID))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set all provided headers
	headers := map[string]string{
		"Accept-Encoding":           "gzip",
		"Connection":                "Keep-Alive",
		"Content-Length":            "136",
		"Content-Type":              "application/x-www-form-urlencoded; charset=UTF-8",
		"Cookie":                    "install_id=7257125026272233222; ttreq=1$aebb5729d2e8c4444a5570cd591c5a6d7829d9c6; d_ticket=8562e6f17992110caa95e548b0b87d2e6841b; passport_csrf_token=88207fccd9dce191a9cb867c010b382c; passport_csrf_token_default=88207fccd9dce191a9cb867c010b382c; multi_sids=7264508143047246853%3Adccefc851f7e801b32d3f0c7e69df1c1%7C7266394608997123077%3Ac4232242555592db4f8b6193e52494ad; odin_tt=51bb494b4a2940342b1413e76780f594c4c3a44d9286e185cc29105bb1c9ef3aa8fa8ad99076bfec33117c025335f7842816b46d597b4f838994f759c41bf66dec725b663ce6296c1284fa4251783420; cmpl_token=AgQQAPOCF-RPsLTuWQ1OcV08xly8AxNP_4QqYNPGYA; uid_tt=19da9371d85993667a66b6ee591d6e844db0af1c48c58de57fc9328acec36b55; uid_tt_ss=19da9371d85993667a66b6ee591d6e844db0af1c48c58de57fc9328acec36b55; sid_tt=dccefc851f7e801b32d3f0c7e69df1c1; sessionid=dccefc851f7e801b32d3f0c7e69df1c1; sessionid_ss=dccefc851f7e801b32d3f0c7e69df1c1; store-idc=maliva; store-country-code=th; store-country-code-src=uid; tt-target-idc=alisg; tt-target-idc-sign=rI-fTXfepuLfspdI1R_MAc6AdSjdviZTMWB7BDlUKymbiarS3bNvf-zGfUX46C9zyZZIOrPwoZRZ-x5qw8ZbF-PdZlBMUEbGWw9qvCgrqhXE3N5tDxg7iu2idlABZlj4sFi2RqGqNTV3jqQJfeQQeF3evAuWQ2UHvtdjs49I-ayZDDg9Q8yBKKNgNfr-juFOKyVgL50qESxY89Nh71uDxDl5NbLKf5m5nYoyv6vonSzrkMbp3GGE_sxeU3ps-9EhRt7HyY3BhkT9oxce6tjKmStWvA61BCHW9TGZblDBKSb85q3d3xZcesUo5HmQV9xCYVCWAxZNc5dNWf9VG-diS7NNzg2FSJaZy_Xvuya8RQcpu48bWQl57CpDuJLB0sI4iJORpCQFTxflD9EliE7-c6_AT5BUSPaqfHUd8ETGbjtfeBi-_XFEXcOEjoUWCqkOMsiKed5EKW6ot8l7zwA-uMJZ7q5Sg54GeTeppWQooHPvTXULaCTX6yL2cLlx_U-n; sid_guard=dccefc851f7e801b32d3f0c7e69df1c1%7C1700302213%7C15552000%7CThu%2C+16-May-2024+10%3A10%3A13+GMT; msToken=-CeLn4nh3XqgFhgHnmFFhIzSBgvhVIyTBS-E0U_6LZC6supTk3hMx_PYi5X-Skr2Fc2BU1wQKH3ECj2aSBdcw2wOMx1xdU1-PElbKtza6Qrbs7CBf50lkSs=",
		"Host":                      "webcast16-normal-c-alisg.tiktokv.com",
		"multi_login":               "1",
		"passport-sdk-version":      "19",
		"sdk-version":               "2",
		"User-Agent":                "com.zhiliaoapp.musically/2023004050 (Linux; U; Android 9; en; SM-G955N; Build/NRD90M.G955NKSU1AQDC;tt-ok/3.12.13.1)",
		"X-Argus":                   "FdUWGDgrq3CHHu6st/wqI7K1PbKC+x2gR3HFVvNhG/x9gxrKpYsnxX8FcjCw9d4zuMDYKH76XvXO+0DuZ2lLpsU8FXNmkp3TA99Yiy+ImjXYY6YTkGP3wNWpBkHjx4Vmdk5g8lkcSSDqs7cYwRnwMEO06dsx54i1CKFNTj9VuEzrifStl9ls3wv4EMlk2w1piNdjhM45MUPDN8c9rnlmFIoTa6FwhtkLBrZ8BBPSUbnCgFurTAjswlZ+ciVuBjmJv1I8vY8qvD4hErf+NDkvd/JdTeB1u/LZfpOebSRn61HqGOK3TDy50Eb8sHMe+VKS9ErjZIE2z8/UuUD+LiHwiu539wHYuSnKO5vWIJiR1Us2ahXs2IWP9R4I9xATzPmwaPgvJtgKDFGTjVkpKcJaBUb/bNkxAAUS9OHFZRdesFLIblu4uP+InHgPA6YVFqLygc2DWwICd2m4xM0zvrN0dZOceWuU58qqkaHGFb7/GIcSXBG8tO7u6UaFP4LaEOY6zmckm0NuQazQbmDuvHRIwIqSa3pEBg20LIK8nVwBCXU//tGpNAnfFAHdujE85JrOou5wjvR69pKYOrVUxwV3QG1dlR1YFOqJokKuIONGhG/7P/4f90o90wQn4UYY/BeBlF4=",
		"x-bd-client-key":           "#LHAdZUfSrepbn/A1hg4pugpSJWA8t6Ch9kDorergMMh28hdOz8GFdWWX5MX3SXtVjcRgNomGklyuxLw5",
		"x-bd-kmsv":                 "0",
		"X-Gorgon":                  "840480e94000adcb03f5bdb764c4a3af87fc7616dc0f9f6a9274",
		"X-Khronos":                 "1700302351",
		"X-Ladon":                   "fVRSUvUdd870k/asQ2qffg7gyoUMAjhdlwgG0YLTLOVW+1Fm",
		"X-SS-REQ-TICKET":           "1700302354017",
		"X-SS-STUB":                 "2D1B1107E185F5C6E3EE3AE921AEF09A",
		"x-tt-dm-status":            "login=1;ct=1;rt=1",
		"x-tt-multi-sids":           "7264508143047246853%3Adccefc851f7e801b32d3f0c7e69df1c1%7C7266394608997123077%3Ac4232242555592db4f8b6193e52494ad",
		"x-tt-store-region":         "th",
		"x-tt-store-region-src":     "uid",
		"X-Tt-Token":                "03dccefc851f7e801b32d3f0c7e69df1c100b9171fdd2e5f4faca7e09f1297d66a8bc60374e992648623079232276125e022b05871379c5e11bfd13e0a736f399d916aaf2a0648b5a4739d7245d30922a04d3c23ab6dc16623afbd881f55503b468d9-CkBmZjQ0NDg3ODg5ZTM5Zjg4YzU1MDBhYjZjNmJiMTI0MzJhZTljOGI2MDc0MjFjYzhlZGFjNTY1YTNiZjZlM2U0-2.0.0",
		"x-vc-bdturing-sdk-version": "2.3.2.i18n",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Process the response as needed
	fmt.Println("Response Body:", string(body))
}

func main() {
	var roomID string
	fmt.Print("Enter Room ID: ")
	_, err := fmt.Scanln(&roomID)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	makeRequest(roomID)
}
