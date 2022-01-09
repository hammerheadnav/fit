// Generated by reader_util_test.go.
// Only edit to bootstrap new entries or change existing entries.

package fit_test

var decodeTestFiles = [...]struct {
	folder      string
	name        string
	wantErr     bool
	fingerprint uint64
	compress    bool
	dopts       testingDecodeOpts
	skipEncode  bool
	encodeNote  string
}{
	{
		"me",
		"activity-small-fenix2-run.fit",
		false,
		16272604713108132935,
		true,
		tdoAllWithDiscardLogger,
		true,
		"Decode mismatch due to unknown fields",
	},
	{
		"fitsdk",
		"Activity.fit",
		false,
		5393063379197673570,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"MonitoringFile.fit",
		false,
		11936585269915402423,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"fitsdk",
		"Settings.fit",
		false,
		18422634047426156243,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WeightScaleMultiUser.fit",
		false,
		15668939108214135507,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutCustomTargetValues.fit",
		false,
		14786462533817802322,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutIndividualSteps.fit",
		false,
		5594723163894392697,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutRepeatGreaterThanStep.fit",
		false,
		17460594330644784595,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutRepeatSteps.fit",
		false,
		4256650519608265147,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WeightScaleSingleUser.fit",
		false,
		16394851171432919279,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"python-fitparse",
		"garmin-edge-500-activitiy.fit",
		false,
		17514013668470756651,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"python-fitparse",
		"sample-activity-indoor-trainer.fit",
		false,
		945649812206588852,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"python-fitparse",
		"compressed-speed-distance.fit",
		false,
		0,
		false,
		tdoNone,
		false,
		"",
	},
	{
		"python-fitparse",
		"antfs-dump.63.fit",
		false,
		6282273622209975218,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"sram",
		"Settings.fit",
		false,
		5866657363356029809,
		true,
		tdoNone,
		true,
		"Fails due to encoder using profile length for arrays (#37)",
	},
	{
		"sram",
		"Settings2.fit",
		false,
		15709312684722569429,
		true,
		tdoNone,
		true,
		"Fails due to encoder using profile length for arrays (#37)",
	},
	{
		"dcrainmaker",
		"Edge810-Vector-2013-08-16-15-35-10.fit",
		false,
		12420128971150793206,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"misc",
		"2013-02-06-12-11-14.fit",
		false,
		11959686082894445424,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"misc",
		"2015-10-13-08-43-15.fit",
		false,
		16776362073923423348,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"bpg",
		"garmin.fit",
		false,
		11468599866908951097,
		true,
		tdoNone,
		false,
		"Has a definition message with number of fields >85",
	},
	{
		"corrupt",
		"activity-filecrc.fit",
		true,
		13015127050946751954,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"corrupt",
		"activity-unexpected-eof.fit",
		true,
		5112456444205297678,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"misc",
		"0134902991.fit",
		false,
		1269717752691992296,
		true,
		tdoNone,
		true,
		"Contains developer data fields",
	},
	{
		"misc",
		"mornindew-broken.fit",
		false,
		864390381975294256,
		true,
		tdoNone,
		true,
		"Contains developer data fields",
	},
	{
		"fitsdk",
		"DeveloperData.fit",
		false,
		7735802126653373100,
		true,
		tdoNone,
		true,
		"Contains developer data fields",
	},
}
