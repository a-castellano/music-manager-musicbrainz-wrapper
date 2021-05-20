// +build integration_tests unit_tests

package albums

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockOneReleaseGroup struct {
}

func (m MockOneReleaseGroup) SearchReleaseGroups(req *http.Request) (*http.Response, error) {

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{"created":"2021-05-18T17:40:07.762Z","count":1,"offset":0,"release-groups":[{"id":"4a9421e6-9def-4616-82bf-674d5a5c7a29","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":100,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Hexndeifl","first-release-date":"2021-02-15","primary-type":"Album","artist-credit":[{"name":"Gråinheim","artist":{"id":"c46ed989-cb1c-48fb-848d-0078784fc3f4","name":"Gråinheim","sort-name":"Gråinheim"}}],"releases":[{"id":"cb36e0ab-6634-4d85-a84e-8be89924811f","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Hexndeifl","status":"Official"}]}]}
	`))}}}

	response, responseError := client.Do(req)

	return response, responseError
}

func (m MockOneReleaseGroup) GetReleases(req *http.Request) (*http.Response, error) {

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{"releases":[{"date":"2021-02-15","country":"XW","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","text-representation":{"script":"Latn","language":"eng"},"title":"Hexndeifl","disambiguation":"","packaging-id":"119eba76-b343-3e02-a292-f0f00644bb9b","quality":"normal","packaging":"None","release-events":[{"area":{"name":"[Worldwide]","type-id":null,"sort-name":"[Worldwide]","iso-3166-1-codes":["XW"],"type":null,"id":"525d4e18-3d00-31b9-a58b-a146a916de8f","disambiguation":""},"date":"2021-02-15"}],"id":"cb36e0ab-6634-4d85-a84e-8be89924811f","barcode":null,"status":"Official"}],"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","disambiguation":"","title":"Hexndeifl","secondary-types":[],"primary-type":"Album","id":"4a9421e6-9def-4616-82bf-674d5a5c7a29","first-release-date":"2021-02-15","secondary-type-ids":[]}
	`))}}}

	response, responseError := client.Do(req)

	return response, responseError
}

func (m MockOneReleaseGroup) GetReleaseInfo(req *http.Request) (*http.Response, error) {

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{"date":"2021-02-15","asin":null,"country":"XW","text-representation":{"script":"Latn","language":"eng"},"status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Hexndeifl","disambiguation":"","packaging-id":"119eba76-b343-3e02-a292-f0f00644bb9b","quality":"normal","media":[{"format":"Digital Media","title":"","format-id":"907a28d9-b3b2-3ef6-89a8-7b18d91d4794","track-count":6,"tracks":[{"length":336562,"position":1,"number":"1","title":"Transilvania","id":"befa427c-5f46-4cb8-84aa-ce8ddbbafbe1","recording":{"length":336562,"video":false,"first-release-date":"2021-02-15","id":"60ab5ee3-c283-467c-ace3-d07927ae166c","disambiguation":"","title":"Transilvania"}},{"number":"2","length":258088,"position":2,"title":"Vlad Draculea","id":"8bd83c92-9d96-41fa-8478-027d5c08fba3","recording":{"length":258088,"video":false,"first-release-date":"2021-02-15","id":"01b07b1d-6cff-4cbb-9c15-79afab66caf6","title":"Vlad Draculea","disambiguation":""}},{"number":"3","position":3,"length":333366,"recording":{"disambiguation":"","title":"Finstersucht","id":"baecb7a6-e05b-48b8-ae3b-300750c3f58b","first-release-date":"2021-02-15","video":false,"length":333366},"id":"ade87fb9-6f87-4635-9736-0f3b4a416550","title":"Finstersucht"},{"recording":{"length":327267,"first-release-date":"2021-02-15","video":false,"id":"651d8ac5-c710-4fbc-965c-52e693d1f75d","disambiguation":"","title":"Aetherbrand"},"id":"74ac1061-52cf-4c53-b84c-9da690c7945e","title":"Aetherbrand","number":"4","position":4,"length":327267},{"number":"5","length":393438,"position":5,"title":"Hexndeifl","id":"17660646-6191-40f7-b5c5-1b39bd4a9333","recording":{"title":"Hexndeifl","disambiguation":"","id":"478c74ce-5c6e-4a3e-983c-b3f37b998b64","first-release-date":"2021-02-15","video":false,"length":393438}},{"number":"6","length":156763,"position":6,"title":"Ausklang","id":"2246039e-315a-4ce4-bd86-6fc81338f9a5","recording":{"id":"5092d660-9fdf-485e-94ce-272fdcbf7afa","disambiguation":"","title":"Ausklang","length":156763,"video":false,"first-release-date":"2021-02-15"}}],"track-offset":0,"position":1}],"packaging":"None","release-events":[{"area":{"type":null,"disambiguation":"","id":"525d4e18-3d00-31b9-a58b-a146a916de8f","iso-3166-1-codes":["XW"],"sort-name":"[Worldwide]","type-id":null,"name":"[Worldwide]"},"date":"2021-02-15"}],"barcode":null,"id":"cb36e0ab-6634-4d85-a84e-8be89924811f","status":"Official","cover-art-archive":{"front":true,"artwork":true,"back":false,"darkened":false,"count":1}}
	`))}}}

	response, responseError := client.Do(req)

	return response, responseError
}

type MockTroReleaseGroups struct {
}

func (m MockTroReleaseGroups) SearchReleaseGroups(req *http.Request) (*http.Response, error) {

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{"created":"2021-05-20T19:17:10.483Z","count":546170,"offset":0,"release-groups":[{"id":"ba62217c-c12f-4eae-859e-5202936914ff","type-id":"0c60f497-ff81-3818-befd-abfc84a4858b","score":100,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":4,"title":"In Times Before the Light","disambiguation":"remix version","first-release-date":"2002-09-03","primary-type":"Album","secondary-types":["Remix"],"secondary-type-ids":["0c60f497-ff81-3818-befd-abfc84a4858b"],"artist-credit":[{"name":"The Kovenant","artist":{"id":"cc6fed9a-a14b-46e1-a5b0-97f740e66e6a","name":"The Kovenant","sort-name":"Kovenant, The","disambiguation":"Norwegian black metal","aliases":[{"sort-name":"Kovenant, The","type-id":"894afba6-2816-3c24-8072-eadb66bd04bc","name":"The Kovenant","locale":null,"type":"Artist name","primary":null,"begin-date":"1999","end-date":null},{"sort-name":"Covenant","type-id":"894afba6-2816-3c24-8072-eadb66bd04bc","name":"Covenant","locale":null,"type":"Artist name","primary":null,"begin-date":"1993","end-date":"1999"},{"sort-name":"Kovenant","type-id":"1937e404-b981-3cb7-8151-4c86ebfc8d8e","name":"Kovenant","locale":null,"type":"Search hint","primary":null,"begin-date":null,"end-date":null}]}}],"releases":[{"id":"3bbc840c-4bde-3b97-99aa-be8971838848","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times Before the Light","status":"Official"},{"id":"e6377520-fe08-3a37-905f-4336868165ae","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times Before the Light","status":"Official"},{"id":"94d4e9ab-42c0-4990-8fda-458e2b8a09ba","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times Before the Light","status":"Official"},{"id":"08c0d4e3-cbdd-44ad-8567-287351a11795","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times Before the Light","status":"Official"}],"tags":[{"count":1,"name":"rock"},{"count":1,"name":"electronic"},{"count":1,"name":"black metal"},{"count":1,"name":"electro"},{"count":1,"name":"industrial black metal"}]},{"id":"422496f0-d2b5-30a6-806e-cf5b7e70c28f","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":100,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":4,"title":"In Times Before the Light","disambiguation":"original version","first-release-date":"1997","primary-type":"Album","artist-credit":[{"name":"Covenant","artist":{"id":"cc6fed9a-a14b-46e1-a5b0-97f740e66e6a","name":"The Kovenant","sort-name":"Kovenant, The","disambiguation":"Norwegian black metal","aliases":[{"sort-name":"Kovenant, The","type-id":"894afba6-2816-3c24-8072-eadb66bd04bc","name":"The Kovenant","locale":null,"type":"Artist name","primary":null,"begin-date":"1999","end-date":null},{"sort-name":"Covenant","type-id":"894afba6-2816-3c24-8072-eadb66bd04bc","name":"Covenant","locale":null,"type":"Artist name","primary":null,"begin-date":"1993","end-date":"1999"},{"sort-name":"Kovenant","type-id":"1937e404-b981-3cb7-8151-4c86ebfc8d8e","name":"Kovenant","locale":null,"type":"Search hint","primary":null,"begin-date":null,"end-date":null}]}}],"releases":[{"id":"aa08eefc-6346-4976-912c-fa8289a07cd3","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times Before the Light 1995","status":"Official"},{"id":"a869dc53-900b-49b6-b7d2-0b737140d8fb","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times Before the Light","status":"Official"},{"id":"0ea00f24-b806-4d65-85dd-f85fe898bdf0","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times Before the Light 1995","status":"Official"},{"id":"44623667-5f57-4ce2-a453-6c2cf0c954a8","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times Before the Light","status":"Official"}],"tags":[{"count":1,"name":"rock"},{"count":1,"name":"melodic death metal"},{"count":0,"name":"industrial metal"},{"count":1,"name":"black metal"},{"count":1,"name":"melodic black metal"},{"count":3,"name":"symphonic black metal"}]},{"id":"bbbe9f55-0dca-49f5-b946-fb6bfb82a901","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":91,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"In Times Before the Light 1995","first-release-date":"2007-01-26","primary-type":"Album","artist-credit":[{"name":"The Kovenant","artist":{"id":"cc6fed9a-a14b-46e1-a5b0-97f740e66e6a","name":"The Kovenant","sort-name":"Kovenant, The","disambiguation":"Norwegian black metal","aliases":[{"sort-name":"Kovenant, The","type-id":"894afba6-2816-3c24-8072-eadb66bd04bc","name":"The Kovenant","locale":null,"type":"Artist name","primary":null,"begin-date":"1999","end-date":null},{"sort-name":"Covenant","type-id":"894afba6-2816-3c24-8072-eadb66bd04bc","name":"Covenant","locale":null,"type":"Artist name","primary":null,"begin-date":"1993","end-date":"1999"},{"sort-name":"Kovenant","type-id":"1937e404-b981-3cb7-8151-4c86ebfc8d8e","name":"Kovenant","locale":null,"type":"Search hint","primary":null,"begin-date":null,"end-date":null}]}}],"releases":[{"id":"5b4e1f75-9bb7-4537-b504-4a727914b721","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times Before the Light 1995","status":"Official"}]},{"id":"2682f8e1-876e-478f-80b8-d2cd74865221","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":36,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before the Light","first-release-date":"2018-03-08","primary-type":"Album","artist-credit":[{"name":"Tape Loop Orchestra","artist":{"id":"35e6d0b9-ec64-46ec-ad3a-931cbeea7474","name":"Tape Loop Orchestra","sort-name":"Tape Loop Orchestra"}}],"releases":[{"id":"3bb12bb8-21ad-4af6-92b6-245e56a20817","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before the Light","status":"Official"}],"tags":[{"count":1,"name":"ambient"}]},{"id":"be3ffb8b-f091-46fb-a51a-acdf8ba4f98f","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":36,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before the Light","first-release-date":"2004","primary-type":"Album","secondary-types":["Demo"],"secondary-type-ids":["81598169-0d6c-3bce-b4be-866fa658eda3"],"artist-credit":[{"name":"Apocalypse","artist":{"id":"696021b7-4803-4297-91e4-6115e6ce5e76","name":"Apocalypse","sort-name":"Apocalypse","disambiguation":"Manchester, New Hampshire"}}],"releases":[{"id":"c48de7ad-ff45-4d31-a105-248f37d1d1b5","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before the Light","status":"Official"}]},{"id":"23b8e49e-a517-3c99-bfec-50d4eaccc81b","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":36,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before the Light","first-release-date":"2001","primary-type":"Album","artist-credit":[{"name":"Ketil Bjørnstad","artist":{"id":"9a141277-3edd-4aac-931b-10030410bfd4","name":"Ketil Bjørnstad","sort-name":"Bjørnstad, Ketil","aliases":[{"sort-name":"Ketil Bjornstad","name":"Ketil Bjornstad","locale":null,"type":null,"primary":null,"begin-date":null,"end-date":null}]}}],"releases":[{"id":"8f84b096-a412-43ac-a11e-0650fbcc468c","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before the Light","status":"Official"}],"tags":[{"count":1,"name":"jazz"}]},{"id":"64cc7f52-833a-4a03-a19c-f1726c843890","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":36,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":5,"title":"before light","first-release-date":"2015-09-16","primary-type":"Album","artist-credit":[{"name":"keeno","artist":{"id":"a746be50-ec09-4591-9cb1-e105918a3c25","name":"keeno","sort-name":"keeno","disambiguation":"Vocaloid producer"}}],"releases":[{"id":"5e462eb4-fdab-49ed-b7ee-2e7020460d10","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"before light","status":"Official"},{"id":"33a872a0-189a-42d4-89c7-b6ff0968dd17","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"before light","status":"Official"},{"id":"75424687-6916-4924-90d3-0c7b710896e6","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"before light","status":"Official"},{"id":"0342811b-3964-42f5-a730-6e46a1a35702","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"before light オリジナル歌ってみたCD","status":"Official"},{"id":"f2d53da4-61c9-4672-84a9-2af34e429d37","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"before light","status":"Official"}]},{"id":"7e00cdc4-d51f-4f89-adc5-b9d1ffa7c9a3","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":35,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before First Light","first-release-date":"2015-09-18","primary-type":"Album","artist-credit":[{"name":"The Oscillators","artist":{"id":"d1cbdf7d-a2f7-4c02-ab71-e38edfa3afca","name":"The Oscillators","sort-name":"The Oscillators","disambiguation":"Australian Drum & Bass"}}],"releases":[{"id":"2d865f45-e583-4193-8a0c-6cc0d23fa1bc","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before First Light","status":"Official"}]},{"id":"1d7a6226-57fb-4e98-af34-275993dbcdc1","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":33,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before The Night Fall","first-release-date":"2014-03-23","primary-type":"Album","artist-credit":[{"name":"Forgotten Times","artist":{"id":"09cd5e94-feb8-461a-b0e6-23a21a03df1b","name":"Forgotten Times","sort-name":"Forgotten Times"}}],"releases":[{"id":"91149011-66fb-4396-bfd8-16628b82c7b4","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before The Night Fall","status":"Official"}],"tags":[{"count":1,"name":"ambient"},{"count":1,"name":"dark ambient"},{"count":1,"name":"medieval"},{"count":1,"name":"dungeon synth"}]},{"id":"554e3dd4-52cb-40de-8de0-a16d8d6cfd45","type-id":"6d0c5bf6-7a33-3420-a519-44fc63eedebf","score":32,"primary-type-id":"6d0c5bf6-7a33-3420-a519-44fc63eedebf","count":1,"title":"Before What I've Created","first-release-date":"2010-09-12","primary-type":"EP","artist-credit":[{"name":"Lost in the Light","artist":{"id":"4204d6ea-59a8-4b8f-8e10-86fe781103b7","name":"Lost in the Light","sort-name":"Lost in the Light"}}],"releases":[{"id":"3e632b23-fbcc-4caa-b5ec-06a4772c8ef9","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before What I've Created","status":"Official"}]},{"id":"01d69863-bd80-44e8-8ece-f7ce863ed306","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":32,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Leave Before The Light","first-release-date":"2019","primary-type":"Album","artist-credit":[{"name":"AT THE SUN","artist":{"id":"35ecc5d7-6519-4f7a-bee4-430bc5ba6b8f","name":"AT THE SUN","sort-name":"AT THE SUN"}}],"releases":[{"id":"13dd201e-e76b-4a5e-a9df-7e5390323476","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Leave Before The Light","status":"Official"}]},{"id":"fea002ad-a19b-3f6c-9563-1937a905095b","type-id":"d6038452-8ee0-3f68-affc-2de9a1ede0b9","score":32,"primary-type-id":"d6038452-8ee0-3f68-affc-2de9a1ede0b9","count":1,"title":"Before The Light Goes","primary-type":"Single","artist-credit":[{"name":"The Fallout Trust","artist":{"id":"993e59e2-2b88-49d0-9a4e-f0589a78c2b1","name":"The Fallout Trust","sort-name":"Fallout Trust, The","aliases":[{"sort-name":"The Fall Out Trust","name":"The Fall Out Trust","locale":null,"type":null,"primary":null,"begin-date":null,"end-date":null}]}}],"releases":[{"id":"988f962f-9b60-48ba-92d7-c14960d6651a","status-id":"518ffc83-5cde-34df-8627-81bff5093d92","title":"Before The Light Goes","status":"Promotion"}]},{"id":"0d86ff06-1463-3f39-9cc2-7b7a5befbd74","type-id":"d6038452-8ee0-3f68-affc-2de9a1ede0b9","score":32,"primary-type-id":"d6038452-8ee0-3f68-affc-2de9a1ede0b9","count":2,"title":"Before the Light Goes","first-release-date":"2005-10-20","primary-type":"Single","artist-credit":[{"name":"The Fallout Trust","artist":{"id":"993e59e2-2b88-49d0-9a4e-f0589a78c2b1","name":"The Fallout Trust","sort-name":"Fallout Trust, The","aliases":[{"sort-name":"The Fall Out Trust","name":"The Fall Out Trust","locale":null,"type":null,"primary":null,"begin-date":null,"end-date":null}]}}],"releases":[{"id":"5296d575-3ccd-4401-9bff-c58ffaaae9d8","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before the Light Goes","status":"Official"},{"id":"b05bbb0f-b5ad-468e-81a9-9dbcc31eb807","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before the Light Goes","status":"Official"}],"tags":[{"count":1,"name":"rock"},{"count":1,"name":"indie rock"}]},{"id":"84340e12-617b-4eb9-b1b8-dba88f265ad9","type-id":"6d0c5bf6-7a33-3420-a519-44fc63eedebf","score":32,"primary-type-id":"6d0c5bf6-7a33-3420-a519-44fc63eedebf","count":1,"title":"Before the Dimming Light","first-release-date":"2012-07-06","primary-type":"EP","artist-credit":[{"name":"Advent Sorrow","artist":{"id":"bc5a782a-659d-4597-bacd-8305b88c2a89","name":"Advent Sorrow","sort-name":"Advent Sorrow"}}],"releases":[{"id":"6e80e411-821f-4096-b9bb-3d01d7d0c926","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before the Dimming Light","status":"Official"}],"tags":[{"count":1,"name":"rock"},{"count":1,"name":"black metal"}]},{"id":"aff5bc9c-7364-4304-839a-fe726d129ce3","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":32,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"33MINUTES BEFORE THE LIGHT","first-release-date":"2013-03-06","primary-type":"Album","artist-credit":[{"name":"Livingstone Daisy","artist":{"id":"936af010-e8e6-4c0d-8778-2ac2fa177927","name":"Livingstone Daisy","sort-name":"Livingstone Daisy"}}],"releases":[{"id":"ca4fb451-d0fb-49be-a813-0616eeb2d4eb","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"33MINUTES BEFORE THE LIGHT","status":"Official"}]},{"id":"048e089f-8fb5-4fd7-a61f-83c27cc7fa97","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":32,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before the First Light","first-release-date":"2010","primary-type":"Album","artist-credit":[{"name":"Black Ice","artist":{"id":"5bb15af5-7182-4c1a-a46a-28b88c19bc6c","name":"Black Ice","sort-name":"Black Ice","disambiguation":"Nowave / Deathrock band from California","aliases":[{"sort-name":"Blackice","name":"Blackice","locale":null,"type":null,"primary":null,"begin-date":null,"end-date":null}]}}],"releases":[{"id":"bb2f3612-0707-462c-99e0-dd5f8ae8700c","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before the First Light","status":"Official"}]},{"id":"312a7d4f-7557-4dda-9d6f-0f6b4737aed8","type-id":"d6038452-8ee0-3f68-affc-2de9a1ede0b9","score":32,"primary-type-id":"d6038452-8ee0-3f68-affc-2de9a1ede0b9","count":1,"title":"Banisher in Times of Light","first-release-date":"2011-01-17","primary-type":"Single","artist-credit":[{"name":"Dies Ater","artist":{"id":"3bcdb124-8bfb-4c55-87e9-cf7ab76d3d33","name":"Dies Ater","sort-name":"Dies Ater"}}],"releases":[{"id":"7153403a-ed66-4b05-be9b-1c6dfc52f9e9","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Banisher in Times of Light","status":"Official"}],"tags":[{"count":1,"name":"metal"},{"count":1,"name":"black metal"}]},{"id":"78877cac-fcbe-32ee-b3ff-20ac05846219","type-id":"6fd474e2-6b58-3102-9d17-d6f7eb7da0a0","score":31,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before First Light","first-release-date":"2003-03-17","primary-type":"Album","secondary-types":["Live"],"secondary-type-ids":["6fd474e2-6b58-3102-9d17-d6f7eb7da0a0"],"artist-credit":[{"name":"Marillion","artist":{"id":"1932f5b6-0b7b-4050-b1df-833ca89e5f44","name":"Marillion","sort-name":"Marillion","disambiguation":"British progressive rock band","aliases":[{"sort-name":"Trios Marillos, Los","name":"Los Trios Marillos","locale":null,"type":null,"primary":null,"begin-date":null,"end-date":null},{"sort-name":"Low Fat Yoghurts","name":"Low Fat Yoghurts","locale":null,"type":null,"primary":null,"begin-date":null,"end-date":null},{"sort-name":"Skyline Drifters","name":"Skyline Drifters","locale":null,"type":null,"primary":null,"begin-date":null,"end-date":null}]}}],"releases":[{"id":"f42dc7ff-359e-4f72-8d7a-8d405d56ec43","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before First Light","status":"Official"}]},{"id":"a96aec31-e205-425f-90f7-efa778915911","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":31,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before daylight","primary-type":"Album","artist-credit":[{"name":"Neal Black & The Healers","artist":{"id":"a51289c7-47cf-4d7c-8455-00b372d16309","name":"Neal Black & The Healers","sort-name":"Neal Black & Healers, The"}}],"releases":[{"id":"35eb885b-192d-4e64-8712-34a5cb449d0c","title":"Before daylight"}]},{"id":"7aadf967-b43e-4df4-999a-1b0e0c723845","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":31,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before the Twilight","first-release-date":"2009","primary-type":"Album","artist-credit":[{"name":"JustMe","artist":{"id":"0ec1992a-605d-409d-838c-bbbf88323046","name":"JustMe","sort-name":"JustMe","disambiguation":"US rapper/producer Justin Long","aliases":[{"sort-name":"Long, Justin","type-id":"d4dcd0c0-b341-3612-a332-c0ce797b25cf","name":"Justin Long","locale":null,"type":"Legal name","primary":null,"begin-date":null,"end-date":null}]}}],"releases":[{"id":"35a7a8ce-d8c3-4a4f-b405-a278dc5e7224","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before the Twilight","status":"Official"}]},{"id":"7ff8dc92-639b-38cc-86f6-17361d906edb","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":30,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before The Siren","first-release-date":"2006-03-08","primary-type":"Album","artist-credit":[{"name":"Ra:IN","artist":{"id":"ca83a41c-10f4-424f-9c3b-d718c85b04f2","name":"Ra:IN","sort-name":"Ra:IN"}}],"releases":[{"id":"b90d560d-2aa8-4326-8630-ac7fc1508940","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before The Siren","status":"Official"}]},{"id":"be14c190-65ce-426f-b289-588373b56d20","type-id":"d6038452-8ee0-3f68-affc-2de9a1ede0b9","score":30,"primary-type-id":"d6038452-8ee0-3f68-affc-2de9a1ede0b9","count":1,"title":"Dark Times","first-release-date":"2020-12-01","primary-type":"Single","artist-credit":[{"name":"Light Bulb","artist":{"id":"496eb1cc-c8bc-45a6-9fbe-674d5bce11ae","name":"Light Bulb","sort-name":"Light Bulb"}}],"releases":[{"id":"224f838f-7cfa-4033-b40c-e10eb4862769","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Dark Times","status":"Official"}]},{"id":"5d912d79-a86e-4464-bd0c-b37840437b0d","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":29,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":5,"title":"In Times","first-release-date":"2015-03-06","primary-type":"Album","artist-credit":[{"name":"Enslaved","artist":{"id":"29bb04d6-5602-464c-8577-f55a27511b27","name":"Enslaved","sort-name":"Enslaved","disambiguation":"Norwegian black metal band"}}],"releases":[{"id":"992aca60-0b8f-4405-960d-66517379d777","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times","status":"Official"},{"id":"69747a48-1ffc-419a-afe5-e80f8fd362ed","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times","status":"Official"},{"id":"8a188d03-c577-49ea-8bf3-077f9ac512e0","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times","status":"Official"},{"id":"cc4263a4-facc-484f-86c3-f75cc3b68911","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times","status":"Official"},{"id":"88b9aebc-5a10-4abd-a1fd-ee572de160f0","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"In Times","status":"Official"}],"tags":[{"count":1,"name":"rock"},{"count":2,"name":"black metal"}]},{"id":"cb1a7799-9cd5-4bf5-9ac8-e2b422839aa7","type-id":"6d0c5bf6-7a33-3420-a519-44fc63eedebf","score":29,"primary-type-id":"6d0c5bf6-7a33-3420-a519-44fc63eedebf","count":1,"title":"31: Before the Light Fails","first-release-date":"2010-07-01","primary-type":"EP","artist-credit":[{"name":"Antonymes","artist":{"id":"151edebe-f2bb-48c2-9bf0-cf8c54302070","name":"Antonymes","sort-name":"Antonymes"}}],"releases":[{"id":"1712dc96-2ab5-4f13-a70f-b358c24c12ce","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"31: Before the Light Fails","status":"Official"}]},{"id":"9f89c01a-8904-4869-9e5f-3a926dc7a405","type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","score":29,"primary-type-id":"f529b476-6e62-324f-b0aa-1f3e33d313fc","count":1,"title":"Before Turning Off the Light","first-release-date":"2004-07-02","primary-type":"Album","artist-credit":[{"name":"GROUP","artist":{"id":"3fa1352c-15a7-4400-8a0d-413b11e74417","name":"GROUP","sort-name":"Group","disambiguation":"A Japanese instrumental band."}}],"releases":[{"id":"ca55fc1b-1584-4aaf-81c6-52547eb39bdc","status-id":"4e304316-386d-3409-af2e-78857eec5cfe","title":"Before Turning Off the Light","status":"Official"}]}]}
	`))}}}

	response, responseError := client.Do(req)

	return response, responseError
}

func (m MockTroReleaseGroups) GetReleases(req *http.Request) (*http.Response, error) {

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
	`))}}}

	response, responseError := client.Do(req)

	return response, responseError
}

func (m MockTroReleaseGroups) GetReleaseInfo(req *http.Request) (*http.Response, error) {

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
	`))}}}

	response, responseError := client.Do(req)

	return response, responseError
}

func TestSearchReleaseGroupWithNoResults(t *testing.T) {

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{"created":"2021-01-17T11:29:49.260Z","count":0,"offset":0,"release-groups":[]}
	`))}}}

	searchAlbumInfo := SearchAlbumInfo{Client: client}
	_, _, err := SearchAlbum(searchAlbumInfo, "AnyNonExistAlbum")

	if err == nil {
		t.Errorf("TestSearchReleaseGroupWithNoResult should fail.")
	}

	if err.Error() != "No release group was found." {
		t.Errorf("Error should be 'No release group was found.', not '%s'.", err)
	}

}

func TestSearchReleaseGroupWithResultsButNoMatches(t *testing.T) {

	searchAlbumInfo := MockOneReleaseGroup{}
	releaseGroup, releaseGroups, err := getReleaseGroup(searchAlbumInfo, "AnyNonExistAlbum", "AnyNonExistAlbum")

	if err != nil {
		t.Errorf("TestSearchReleaseGroupWithResultsButNoMatches should not fail.")
	}

	if releaseGroup.ID != "" {
		t.Errorf("First ReleaseGroup should have no data, not %s.", releaseGroup.ID)
	}

	if len(releaseGroups) != 0 {
		t.Errorf("releaseGroups should be an empty array.")
	}

}

func TestSearchReleaseGroupWithOnlyOneResult(t *testing.T) {

	searchAlbumInfo := MockOneReleaseGroup{}
	releaseGroup, releaseGroups, err := getReleaseGroup(searchAlbumInfo, "Hexndeifl", "Hexndeifl")

	if err != nil {
		t.Errorf("TestSearchReleaseGroupWithOnlyOneResult should not fail.")
	}

	if releaseGroup.ID != "4a9421e6-9def-4616-82bf-674d5a5c7a29" {
		t.Errorf("First ReleaseGroup ID should be '4a9421e6-9def-4616-82bf-674d5a5c7a29', not '%s'.", releaseGroup.ID)
	}

	if releaseGroup.Title != "Hexndeifl" {
		t.Errorf("First ReleaseGroup Title should be 'Hexndeifl', not '%s'.", releaseGroup.Title)
	}

	if releaseGroup.ReleaseYear != 2021 {
		t.Errorf("First ReleaseGroup ReleaseYear should be 2021, not '%d'.", releaseGroup.ReleaseYear)
	}

	if len(releaseGroups) != 0 {
		t.Errorf("releaseGroups should be an empty array.")
	}

}

func TestSearchReleaseGroupWithMoreThanOneResult(t *testing.T) {

	searchAlbumInfo := MockTroReleaseGroups{}
	releaseGroup, releaseGroups, err := getReleaseGroup(searchAlbumInfo, "In Times Before the Light", "In Times Before the Light")

	if err != nil {
		t.Errorf("TestSearchReleaseGroupWithMoreThanOneResult should not fail.")
	}

	if releaseGroup.ID != "ba62217c-c12f-4eae-859e-5202936914ff" {
		t.Errorf("First ReleaseGroup ID should be 'ba62217c-c12f-4eae-859e-5202936914ff', not '%s'.", releaseGroup.ID)
	}

	if releaseGroup.Title != "In Times Before the Light" {
		t.Errorf("First ReleaseGroup Title should be 'In Times Before the Light', not '%s'.", releaseGroup.Title)
	}

	if releaseGroup.ReleaseYear != 2002 {
		t.Errorf("First ReleaseGroup ReleaseYear should be 2002, not '%d'.", releaseGroup.ReleaseYear)
	}

	if len(releaseGroups) != 1 {
		t.Errorf("releaseGroups should be an array with only one element.")
	}

	secondRelease := releaseGroups[0]

	if secondRelease.ID != "422496f0-d2b5-30a6-806e-cf5b7e70c28f" {
		t.Errorf("First ReleaseGroup ID should be '422496f0-d2b5-30a6-806e-cf5b7e70c28f', not '%s'.", secondRelease.ID)
	}

	if secondRelease.Title != "In Times Before the Light" {
		t.Errorf("First ReleaseGroup Title should be 'In Times Before the Light', not '%s'.", secondRelease.Title)
	}

	if secondRelease.ReleaseYear != 1997 {
		t.Errorf("First ReleaseGroup ReleaseYear should be 1997, not '%d'.", secondRelease.ReleaseYear)
	}
}
