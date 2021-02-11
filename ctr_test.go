package drbg_test

import (
	. "gopkg.in/check.v1"
)

func (s *drbgSuite) testCTR16_AES128(c *C, data *testData) {
	s.testCTR(c, 16, data)
}

func (s *drbgSuite) TestCTR16_AES128_0(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "890eb067acf7382eff80b0c73bc872c6"),
		nonce:           decodeHexString(c, "aad471ef3ef1d203"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "a5514ed7095f64f3d0d3a5760394ab42062f373a25072a6ea6bcfd8489e94af6cf18659fea22ed1ca0a9e33f718b115ee536b12809c31b72b08ddd8be1910fa3"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_1(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "c47be8e8219a5a87c94064a512089f2b"),
		nonce:           decodeHexString(c, "f2a23e636aee75c6"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5a1650bb6d6a16f6040591d56abcd5dd3db8772a9c75c44d9fc64d51b733d4a6759bd5a64ec4231a24e662fdd47c82db63b200daf8d098560eb5ba7bf3f9abf7"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_2(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "130b8c3d2d7b6e02c410b4168e122c38"),
		nonce:           decodeHexString(c, "79a674c5b2c51aa9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "57e8a1e578ede1c66879c430df726435d51a369a0fe59a0358d1de352d4280fd7b225f5f386a4fcf12f72794ad0f3757fb25deba3c7512ce4d3733c7ee067043"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_3(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "04d24145287664f6762b5d2a102ac603"),
		nonce:           decodeHexString(c, "ecac63e1217ee335"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "fa5e746decd6801eb7083b6f0e72432e1fd4243104f748d0f19083923b9555688f43146d5acea962da01231d9e5faff0e81f3d394ace3a3454536d726575041f"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_4(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "aede4e613556b1d5a30fce261fbb820c"),
		nonce:           decodeHexString(c, "39acba03c5f10af4"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "232c44b819b88f1aeb83f2034f842d5a00f0301505d2aa69aaecb3cb14bcb15875e0fd60071a80f6262dcebcf41a0e1476d96f409712d828ae313a9d28ec2dee"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_5(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "9f90541c10d4b7c089fe688ea3ef4fc6"),
		nonce:           decodeHexString(c, "1eac1c22036e2b22"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "71af3fdf673404163b06737e0f39915faec21821816c3142e80a503c70cb2edd468c3f03cb1f8a2f9248635b63d7b8f19e9e11caf0ab0c3d2ff9c71321cac154"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_6(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "3b8a181cea8390d5d7a4e6515cf92e3e"),
		nonce:           decodeHexString(c, "37f4d1748714345b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d6fd0ffb10850eb6eb7050e61eac00d472f65cd3d935081477fe44aa85694e120ab5a1ca1fa198af76dfa8d0abdf53e85aa8c87fed0a8c24163943b96d80aafb"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_7(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "854caa2a74f3f43b6abd806d6748ed80"),
		nonce:           decodeHexString(c, "0cbd1372beb62736"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "46534028165ffe2845a972627e67e153510ca1880e6a3ad31fdee71ff240d3278624b31f2d386d7b2228ced624a42e7a3b07480b2323166c18d1ac0f60002ec4"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_8(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "87e1c532997f57a35c286de864bff264"),
		nonce:           decodeHexString(c, "a39e98db6c10787f"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "2c147e24119ad8d4b2ed61c153d050c924ff597515f1173a3df44b2c8428ef890eb9def3e47804b2fd9b357fe13f8a3e10c8670af9df2d6c96fbb2b8cb2dd6b0"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_9(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "2fc623429029c96edef6166247b08cb0"),
		nonce:           decodeHexString(c, "9982663355582788"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5596cb16f3be8552c1e5c164d540cb1faf4bea8733b60a8ad0c40626256548c7aa96d1d2729d26f008731fc39307be5bcd2081c69e314e0c73e3d0fd1d905828"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_10(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "98ab8b4eafab6e536f7845abec137808"),
		nonce:           decodeHexString(c, "dba944c98b311d8e"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "86eed3a9fa53452fb11dba9cac8e4402522928f2705a5e582f4d00eb8fed818e629c72a6a779beb4ed9a239368233cbfcf55685dbf2de34ab58920cfaca4aafe"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_11(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "7dcf4fa731139c5bb6442fed91e89f68"),
		nonce:           decodeHexString(c, "ffd94fe21a808b15"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "8eca20e31c9839b741aaa9bf6ceee224d32126b3196eef3ecd343d414d3233b9fd0ea0ed1bc7700c88cd7c88d3c07613c42cd1f909fed8c6a708d05d6b68fb2e"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_12(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "519c6efeded2a110ac41839a8b8abfce"),
		nonce:           decodeHexString(c, "81d95edc06dde6b3"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "39cdd10e49e03581e3811ddd07d9d0acc34085a12c1c6b87a76386e97032dc01f523f632ec950d0434c25937e432d98554599a855db9adf58b9e04597a21d00d"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_13(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "80c19eeaad1a58379466afae1d80876d"),
		nonce:           decodeHexString(c, "d9a2093f11892b82"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "cec51b985ec97b18ee8ead36157beaf96f129f28280889ec873c27b5627198c585a6ad21ae23a959c9fa49d985af0df4028fdf1f51d82e8f2b3f028853f14e8f"),
	})
}

func (s *drbgSuite) TestCTR16_AES128_14(c *C) {
	s.testCTR16_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "b77bd22efeb771586d516f582192a311"),
		nonce:           decodeHexString(c, "a699f42a4981fefc"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "7ff0abeaffc6ec923cebd910f5937bf14fc52d2a74258388c76c1bc9e635adf175858f0a55877e7e9f5f867d00b9b1363dde46288a6fff5321f469445aad4198"),
	})
}

func (s *drbgSuite) testCTR18_AES128(c *C, data *testData) {
	s.testCTR(c, 16, data)
}

func (s *drbgSuite) TestCTR18_AES128_0(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "e10bc28a0bfddfe93e7f5186e0ca0b3b"),
		nonce:           decodeHexString(c, "9ff477c18673840d"),
		personalization: decodeHexString(c, "c980dedf9882ed4464a674967868f143"),
		expected:        decodeHexString(c, "35b00df6269b6641fd4ccb354d56d851de7a77527e034d60c9e1a9e1525a30ed361fded89d3dccb978d4e7a9e100ebf63062735b52831c6f0a1d3e1bdc5ebc72"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_1(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "ca4b1efa75bd69363873b8f9db4d350e"),
		nonce:           decodeHexString(c, "47bf6c3772fdf7a9"),
		personalization: decodeHexString(c, "ebaa602c4dbe33ff1befbf0a0bc69754"),
		expected:        decodeHexString(c, "59c319791bb1f30ee934ae6e8b1fad1f74ca254568b87f7512f8f2ab4c23010305e170ee75d8cbeb234c7a236e1227db6f7aac3c44b7874b6556744534300c3d"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_2(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "3519491574719ddb1f51b5f59e21ad3d"),
		nonce:           decodeHexString(c, "ef1dbeae79f0898b"),
		personalization: decodeHexString(c, "7fdc21e353249e93df98f29102b9acec"),
		expected:        decodeHexString(c, "192c71857447af82013706b8eb008f8e2c2eb3994a238cf739e17460fe84ef71ca437e8acbb8d9d28807c47a5f034442aaf3c2f3e7debe531bc3596056685848"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_3(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "68edc1c7f04fa22f751b918b016bbfb5"),
		nonce:           decodeHexString(c, "a9b67a3958f202f0"),
		personalization: decodeHexString(c, "9d6fbb020e3d99d195ff469f0ac0b8ca"),
		expected:        decodeHexString(c, "d90289b1b6763a7690fcb5d67c81c203a1cfb08871b94af651c4cd7ec2fa4b9d47c8e5fcc99affc22143b6127018e9e2c0914f9d47a5a30b26e74ddc31a44842"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_4(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "8e6baf81a3f9e732efb965afa59e71ee"),
		nonce:           decodeHexString(c, "9a4036fe7ff225b8"),
		personalization: decodeHexString(c, "c57e42ad7aef5377fc4cd620d631b061"),
		expected:        decodeHexString(c, "7139a2d674eefe54087a2dd5f497f5a1cb444472e364045cb9202bce24581b4b1b803fa5f4a9ac6d6637bdd906af3d8f49b5b80f44cd16ee5b1ae3bf51b598ba"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_5(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "f8992f340f90c8d816b4bd8b435da035"),
		nonce:           decodeHexString(c, "fcf944006f292798"),
		personalization: decodeHexString(c, "5632235e852c074009b497eda438a16f"),
		expected:        decodeHexString(c, "9eab55fa4b2c1cfac63d4de47aa8f59f785ee490da44cee30f9161e93afc461dad0e7f4a064e92efd6008916c22f8364c190467ade12ab4e18b1f39ce18b358a"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_6(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "07438f8353b7b0674218319e29540bca"),
		nonce:           decodeHexString(c, "411477636a500c62"),
		personalization: decodeHexString(c, "cf36f2aa72f2f35e335e65cfd17e2d3d"),
		expected:        decodeHexString(c, "6cfacfe47acf8f61d8af3ed4fbacef5cd741ac2f165e15bdaab1d030c68567837ef3da1ec83da0423ca3424cf293f61f7175fba8e91e5fdc8b3955777361d690"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_7(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "8e93735c34ae0c5ea5213d83a8da8045"),
		nonce:           decodeHexString(c, "2a0b279a9d677838"),
		personalization: decodeHexString(c, "14f45266524a8152857a83316a177585"),
		expected:        decodeHexString(c, "e00cb8636f9602c33021a82a7f9dc9c74ed3435b34897d6efd8c05f177e8556a443bd18c20a557b9bdc417e8998506485091d6bc918e9cbffbb85c620f74c6ec"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_8(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "ec6b87e9e03f44f919e40e1dac02f873"),
		nonce:           decodeHexString(c, "5e9ac5090a5e2897"),
		personalization: decodeHexString(c, "0d9ba462bf166f3b366a2adf55d27645"),
		expected:        decodeHexString(c, "ec37acbddfb041e4249c1cb9833a31fc6fbdaa6658e4b748c0ebf6ab54a3b9c0f62d7c89dbb21d1de13315cbae2ff4f5462491321fbe04b414fba2807ed914ee"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_9(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "9d282f5a0c15a653047310b5627d8fa2"),
		nonce:           decodeHexString(c, "5522c42f4bb076dc"),
		personalization: decodeHexString(c, "846b2e17b062c4c7c561943b5d97565d"),
		expected:        decodeHexString(c, "45f60e1ba981af7f52ef4939c0bb0bdb6fe46f372cc648064bbd3a0a2b85c02391b29792cc8805bb5d453ee290cd1b9c9f8f20b09116af1fd5eceae80d0358d8"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_10(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "83827c08524b54e80f60e186cfce36db"),
		nonce:           decodeHexString(c, "177e5c00f785a6f9"),
		personalization: decodeHexString(c, "7175a2c22eab4d0b1a65fdaded356aeb"),
		expected:        decodeHexString(c, "8424e76c736c0309423cf48dc105feef228dd47bfd8d33801950d2102d5dbafb51e1a85b4a971e4f0b9d52656973db455b286f588ca61d1599c4ec60dd80be98"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_11(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "9f3213a2d6fe5f3dfaf0e4ed292e0f2a"),
		nonce:           decodeHexString(c, "f9c0553b18b21cc3"),
		personalization: decodeHexString(c, "2173d39b3755041911989f89f1e95d3e"),
		expected:        decodeHexString(c, "1045f7b536dab4de015448ac4766188a367bb69397913dfd3d5bbf83f1bbac73232d47644e5a9c169960581e4ad5ed2695360b94f4209c77d7694016b4b397c2"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_12(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "fbc8199074a76f23688636ae3114281e"),
		nonce:           decodeHexString(c, "ebc22e95ddaeb6fc"),
		personalization: decodeHexString(c, "71fbd59f88fc99a9b6070027e73541c2"),
		expected:        decodeHexString(c, "872566a7004798e00697bdb33af1ee7a076e9fd289af83afcb10a0252bb5c106c8dd20d77c859f1401118c24f5a8968f11fa9f7778c2803189c8da87e64945c8"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_13(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "43f4d089060c907f3876051fd5e36f74"),
		nonce:           decodeHexString(c, "4571e6615ba3bbce"),
		personalization: decodeHexString(c, "fc9057749ba437d8fbe9f81d29df5ef1"),
		expected:        decodeHexString(c, "9078f1bc92910ebcf8e31757b624bd23dbcd74f9ca70ff1c6f2b21d859a4e8a4e9963ef132a025b3ae285b43a69698900d4f8a30bb5e99290e4504161997a837"),
	})
}

func (s *drbgSuite) TestCTR18_AES128_14(c *C) {
	s.testCTR18_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "0ee38ea29d7195003e2b942abf1348fc"),
		nonce:           decodeHexString(c, "80a75c77ceff54e4"),
		personalization: decodeHexString(c, "838eaebc982caa1562c3064cce42e51e"),
		expected:        decodeHexString(c, "66291b8685e97c76e2216d708b407022684b28291f3dc71b5d60de143f5fdf0bb07ae9d7524456c4dfd089ea88de86ddfa1d48ca542f00fa586da7a6027d837b"),
	})
}

func (s *drbgSuite) testCTR20_AES128(c *C, data *testData) {
	s.testCTR(c, 16, data)
}

func (s *drbgSuite) TestCTR20_AES128_0(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "2d2ab564202918c4ef5b102dda385a18"),
		nonce:           decodeHexString(c, "259195269ec11af6"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "2c5cd79ed87622a91b8654c8903d852242cd49cb5df2d4b4150584301c59f01fd95a702ac157c84cc15f42c8211335672d8ce1291ef9b1def78149a04fa2697c"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_1(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "a016463dbb499990cbcda45046d8f337"),
		nonce:           decodeHexString(c, "249d02de2dcf3e57"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "35b81fb94241f4c9319a7f16b442495252db4c984860d035f6c45403f974f534fa93b21b6b89441be07c5c29317f868dc9ab5c18377437fadb4d857ee092f923"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_2(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "1e3820c46e50b7cbc58b05aa722427a8"),
		nonce:           decodeHexString(c, "fe0f40014eb3279f"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "18530b9de889d8ad3f41d539796c2a95d682fb026df7ca094af4ab2395292e6fcdb175ba257d61279da4d0fc78fdf8a4eb46a0e2d754f5f8f09ecb157c8a58a1"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_3(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "3b02c5a474679648e2fe090a13be041c"),
		nonce:           decodeHexString(c, "71ed5aa078b83a7b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "e8e2e5d8bd9929e4e1675c8461457cab0633d83f9bd243bc582937c2b961b71c11856f836c37329d3fa86376eec97139104de53260e122cb2b6d824e21e75ce7"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_4(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "050c48ab8a05ff690b62e37d9b21f2fa"),
		nonce:           decodeHexString(c, "f7964875be288353"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "24cb0edcb89eec617613a27fe3e4475affab498d57bc3c4ef844fe19b38413477c1db040f16757f1851b799c1f2669aa9f9c50532eafbe1e2948e338d1ca4bd7"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_5(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "ecd99c0b491cf6a323bc333646e5c56a"),
		nonce:           decodeHexString(c, "b10a1784995cb34c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "6535c48d999a306acb83f7b720b3c39772a496e75a890f6e14d70a824142a12d01dee55271b82c7b027b9468c309fa2d08e91e5e379eb90cf8915cf4cc241b50"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_6(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "34aee3ef35b2a1e31c429725feaa6bfd"),
		nonce:           decodeHexString(c, "e4369b6ada9091ed"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "af9db4df46ac80fecad6c3d1d975d2d0a665f31b4e3b07d962632296b76039ddcab8edde17bad6230131162e59d07dbbd0f0b52e5ca93cf37f12c03042bf0c82"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_7(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "91156469540c785df56a10e16b66107c"),
		nonce:           decodeHexString(c, "7e2b2a7f58f3ef20"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "9a763a4c052725e7ae9d2122a45cd18ea28c58172b87c57edb9daa32f8a26eb98d6b8812095a54cf6282c5e0fc18cd30124415d267c9fc2cdae9beedbb86584e"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_8(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "cd5dbc35d1b4435c4c945bafe4b8fc7a"),
		nonce:           decodeHexString(c, "af4acdfc8119322c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "fe678299f1ade19621880235fa6fe2b50184f81243833ebe68c616a9034e23e0cc873e8ac4d0e365b524a37832b52d97a83862932491d2643a73e46a80b5149c"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_9(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "d52d1159787baab213e87921fd959110"),
		nonce:           decodeHexString(c, "7ab37538ee0245c8"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "f4d9ead3594606ac51942215aa0bd29baa7c281e62d2bcc1a77f57e3596683199d70c009c084eb3654224ed7b07a09b494480a113b33e0366cb2253edbb0cb5e"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_10(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "5110db9d5f4fbdc037550e83c2bd96fe"),
		nonce:           decodeHexString(c, "dee641f9177c9da4"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "05f8cad2d031ce6b527ee043b5056db20ee5ca660a95c29b8f8b2e43efa5c7aa62351233f4f160368c3dad4a74f471bb02153f86c87a140a36e0f122c68ba77d"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_11(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "979065bc1a1c8c8290dde219f0121abe"),
		nonce:           decodeHexString(c, "593a0e1fa98564a3"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "199e68e5b637a27be79453d9953a0e529536374e7f63e8595afe4b5d4e3bb379c6d462f6431244af4336b17f9dfdfeafa6f4fff5004aa65f7d6bb3b88af8e9e2"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_12(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "89a787be64ba3358952f9d602bd8fab4"),
		nonce:           decodeHexString(c, "ab961b68c65892d3"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "6e5dd2a651f2614fca881f5edb05d5eec61275dd05b31fd8a3062fac515960def7b0d4db304f642b91e6401f2a00b0fb693edd7313102f58d73302ea8b5a9d7a"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_13(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "21f02fc75a282d7f87cab05767558e08"),
		nonce:           decodeHexString(c, "3b43d77b0c6471e9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "3b8709cefe270627eb4947f3c7695eed2c85cc1525be253f200a44ad55f7583e1091b880ee16a808e90dc1dd5c0ed3b8d196407661de8061f6a2d78cdc93429d"),
	})
}

func (s *drbgSuite) TestCTR20_AES128_14(c *C) {
	s.testCTR20_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "df5b1cd51045c83801532449d6a87aed"),
		nonce:           decodeHexString(c, "8dd172b81286a238"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "f0172b3a5b42ead6e654d3abc78dda08780ea43223625707f8ab021bef5412db89b089cf751178eb1b97be13f5e83b9f1076d0fc88e6bbd6eef1ef6355e3672b"),
	})
}

func (s *drbgSuite) testCTR22_AES128(c *C, data *testData) {
	s.testCTR(c, 16, data)
}

func (s *drbgSuite) TestCTR22_AES128_0(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "0af13f645902af49e9a7ced6e36a210d"),
		nonce:           decodeHexString(c, "c3bff291a11ac497"),
		personalization: decodeHexString(c, "e8f1d1b4731c4d57d7ead9c2f600fdc6"),
		expected:        decodeHexString(c, "ac6f945a4b9fd3b47c74379eb1f2a7bbedf8eec68efd3c7a6cf68c54ae7a3f7be7280f459c2e0b72afa45cebbebba17c867e9611c896a57d515beb06a7b91f4c"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_1(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "0e113f47f2fc76e83e2d13d572245608"),
		nonce:           decodeHexString(c, "5ff85cc6a534f15a"),
		personalization: decodeHexString(c, "50250668e59de35fde91e08fe18484ab"),
		expected:        decodeHexString(c, "dcc64a966a52d6008dbe07a2484bcaad67b254d6f246e4501d9864b64ad8b7edf10fdbc6ddc414a9b431b058a7ee5ced23f7a6ac7eea0fe6131c9eb7412e68df"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_2(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "16d2a233f0497c751ddd85ef7cd862e6"),
		nonce:           decodeHexString(c, "7e7e1aab1b82675d"),
		personalization: decodeHexString(c, "15f81e40a2493b859cced33cdf7b7590"),
		expected:        decodeHexString(c, "c52097d2d009dce7cbd46d740de54d70c0732f96fac0bd169f3856e50f9e4601672538b23f371f780e61e82f2db99ef3834008aadc26c27855a68696ee812c9c"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_3(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "ea36891e0d6426945a6e1f338c86ca4b"),
		nonce:           decodeHexString(c, "896966600e85d3da"),
		personalization: decodeHexString(c, "b8c701a11e7c008e37c232627c24b6d1"),
		expected:        decodeHexString(c, "446360f16b1e59a82596e61df23af18f5a95b7d8b7dd6172f3575d45251386b04b69785ba44a75eb2a5415713a84365cac81b30df0b48a610977fe6961e7ee8a"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_4(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "3f4138d84d7abd13bbe15e18b292f15f"),
		nonce:           decodeHexString(c, "3462e52e82f28653"),
		personalization: decodeHexString(c, "24c5177084951cf532f2432a1882596f"),
		expected:        decodeHexString(c, "ba2ca1e443c75d55c42065ba91f77db5b8de6c342b65a2c149445979ff263e9d018f564430fe1edb48b403cef23f860ea27146d8511843e0a39ff337ba74d221"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_5(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "a6aef80fab61e27ef611b86f924e976c"),
		nonce:           decodeHexString(c, "74278f5023c2552c"),
		personalization: decodeHexString(c, "a3851679e112948201396d98d6be5c13"),
		expected:        decodeHexString(c, "685726827f3d38427d07dfcfb64a7f95b6f06ef01ced281dce7f41303412361124362ba51ce8d2552af054a20fdac94e5169165b715267638c1e0b9b3452b5b7"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_6(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "0164ae53e533f2f1222d4aa27276814c"),
		nonce:           decodeHexString(c, "2605f589cb1bea9c"),
		personalization: decodeHexString(c, "9de703a5e2bbf2e53d0fe3b573f52acf"),
		expected:        decodeHexString(c, "fd7f64622e87de3b66e19930cf4ccdc3e672f5a4def0daddb37107430ec7f691c6321545103c2a614d7f0f33173bdec420585822cda8c272201922b602f5d8a7"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_7(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "29dc8bad500cce295b9e86b2159c6dcd"),
		nonce:           decodeHexString(c, "b69a5a77694fc4a0"),
		personalization: decodeHexString(c, "c53b13afbf21f17ee9f4dd7c0993e0f9"),
		expected:        decodeHexString(c, "08c97d638b5d82599310325d3b1e6327741cd91887d658b5d815eeed7a72ab86e2fa17e9d013e17a9214d6396ac1481f768830a4d8c4fc392012e9076434939f"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_8(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "b3aee2cf80e3d05aeda0e01cf65ea965"),
		nonce:           decodeHexString(c, "24b58b7c6a99eb2a"),
		personalization: decodeHexString(c, "7344043c84915bcdd6d81f3fe23ba273"),
		expected:        decodeHexString(c, "be353cf832e846444af0015ed05d54b72c9a1a0231f2a4c475611746ef861029e18842c1b76148d2a3eb0174a30cad81b558d0d1e01a8c7dd510f0b56f1fddc3"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_9(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "a525f69165c44f3a8af80857a5e874b0"),
		nonce:           decodeHexString(c, "0c806a40c1dcb7cc"),
		personalization: decodeHexString(c, "4312c2309a1dfe2d9a66f5c18fc40b87"),
		expected:        decodeHexString(c, "d0b46f7e1877db244ff4e0aeb64c4efa3e8c2a8662415cf5ba67f2dc91e82310a6a933c4df6a402f8fae270b229bf254cd35943a13340979376669119e0e5cf8"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_10(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "0608dc6c8f9f1956fbe7d5a7c82fd3b3"),
		nonce:           decodeHexString(c, "615b62a00d6d8d85"),
		personalization: decodeHexString(c, "70dd43c1764318174a2857c420ddf407"),
		expected:        decodeHexString(c, "71f4b5df5bfa5fd52cdc9851a633f77ed90458b5a290b04d2f35130d67a42c1698ac5f9b139a9ecf3590755a204160a3a8f17b77726652c6dc6e9f00966454fd"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_11(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "7b391f35db0e53bf48ee5576b023ff90"),
		nonce:           decodeHexString(c, "7621aad532982b23"),
		personalization: decodeHexString(c, "a6e75778b777d01f006227c400d8ff86"),
		expected:        decodeHexString(c, "d61780b5c781412bf3085eeccee49b99358a183223828c1d8013ace613d89a4504d75ba309e510589a53b472bd5fa1ee2a22392b8265707c15e32935bc8efb4e"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_12(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "099f1f591707d1f9abae7e2b65adabc3"),
		nonce:           decodeHexString(c, "f90d357a901f1135"),
		personalization: decodeHexString(c, "74b8260ced4bbb80417792f193c6f9ac"),
		expected:        decodeHexString(c, "86f506256d29450f1607312739e5a089b6b786fdf52145769cb70ef4eed32109cf62369d7a765955781eae63520275169368257c6d348ecc900a3898778d9a1a"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_13(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "c9059d7c015c64f6e41aa492099f4ef0"),
		nonce:           decodeHexString(c, "3f6e0769cee177c6"),
		personalization: decodeHexString(c, "76c6c6613298d925a6c507ae654d6d9a"),
		expected:        decodeHexString(c, "238111c006e514c05f2ae935a32dbaa0b128daa0736e83d93403776c91e477d0d02c4a24876a2329a0f106038d701feddf0247cc75846dd30108299d840ec4a8"),
	})
}

func (s *drbgSuite) TestCTR22_AES128_14(c *C) {
	s.testCTR22_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "98d13c30a6dae2ca76d7d6679f1ffbc4"),
		nonce:           decodeHexString(c, "490b939eb556ff1b"),
		personalization: decodeHexString(c, "6c10a7c30fb534b5934b1c29808494b9"),
		expected:        decodeHexString(c, "cb7deeea97bdb03f72a03b2d9b5f1b330344486e53e04927c58fb63c971ac98733b0249a1efa85f88bec4818e13110220776e4fde429e6390af466712901189b"),
	})
}

func (s *drbgSuite) testCTR24_AES128(c *C, data *testData) {
	s.testCTR(c, 16, data)
}

func (s *drbgSuite) TestCTR24_AES128_0(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "2e1724db482232a3e61f92c1c266faf8"),
		nonce:           decodeHexString(c, "38aa5590f6bfaa4b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "4438b48a45fb0141e31f0a9624dfe6fcc2f9edc075c0a52bc5fc46d85a966c853feee6af913234b3f9a679f667898dc15a24aaed89f035bfa5da516e435bbad1"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_1(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "22564f77c45b053cdf61433eb96b1d7c"),
		nonce:           decodeHexString(c, "cf73e620f8515203"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "c7908e712c716d1f5ed533e142e72187ea77fb4f516dc31aa10a1e549d85eadb7a4646170464c1f7a752c01a9406be6643ee967d0464b84b6a08b2ed0a7acb07"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_2(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "2a5b012b097926e8f8570ff8692cc5d1"),
		nonce:           decodeHexString(c, "de8e072d1581afe6"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "81f2e2c8585e8dc465d78e7b75b9f6c2bbdcde9475d4250fb49e04c56e30489e24df4858f74ed085cba9f992eb7d13e4e064a745f451cb6edfc37c57f35e8d57"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_3(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "2c6b05d1c86aae86a89e816482983236"),
		nonce:           decodeHexString(c, "3385d929ade996e8"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5256a64429a556d1a91d58999c75b36de7cc01f7463c4e24afd15de0a35dcb5ada2679134f15f4c51dc06b34454d6dcaa1d2511c1d226f232f445276255751e6"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_4(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "48bab2f82d80597f93addb7e1f0f2e72"),
		nonce:           decodeHexString(c, "f5f49018e8559b0a"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "fa323d2dae9674d80bd2f1427c1c79532b2374ed1fb3a13c620605048ac578b3070c6748314e5ed1dbd3eaaa641e505c3d3f59fac25d897bf394dcadb63b7ff9"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_5(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "ad7aff4224e93f323545416a1e5697ce"),
		nonce:           decodeHexString(c, "9cc12ce2ed5e8d1c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "36c72223b9244cdb2c2c0dd5a59706558e2e5a11845cabf9545cd4ad08154a46703ae750b7f0c4f5bb33acd3c381e5ee4ce09916431144e8515fed914d5c5b5e"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_6(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "299c01d3a2f1323df753cf14845e0de5"),
		nonce:           decodeHexString(c, "511e36232a11291d"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "7709fdc6278db44b21d39a194b806e48e7fe3e9ae116e238c205c2c345981ce81f255713597cec2b3ad3910f2b6742aea66404712df8328d2f2d1948413097db"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_7(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "b6ee7779994ba8cd490a03cf6899b135"),
		nonce:           decodeHexString(c, "66eda9b5a54d7ded"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "4e21b48fde082263d76a1034e87a566e1a1c9d2e1bd5c748e30e1d8750f2ff03931c4bfe194d2da4ed1cf1530301e5b1abc4bd2bda7be89284f8c2193688c982"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_8(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "e3ccb991c3e1b3eda0b40b5142ac84d3"),
		nonce:           decodeHexString(c, "998757e00da2b9ef"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d32bc190998f18e9d5509f46022905d011bbec774f0583691d48124379e81d99f0cdd46138bccc47eb773f257a662b798fab275886948936bce72dbd2c6188ac"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_9(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "c34e398041ccee23897f7ca343f35640"),
		nonce:           decodeHexString(c, "a24f8ea0886bf6bf"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5ff0daef3004c503b8098e3b968a8e323319be786c7b742aacee355f1a3c9de75061108e79181361f2e3306af07bcded10e32def1b7bb3e4ebc17096c6933058"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_10(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "fc1fd25ced5bd30a621cd4ae779680ad"),
		nonce:           decodeHexString(c, "af17a9c974b36e6c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "59f99d085749367478684a5ddc8fe181b97a4e67fde5c151c4696d523d7c14b72689a95a5b6092e949dd163bd8f9e45727d2b8a31ed288c8c6229e8be6808ec8"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_11(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "72c0e23d92070a8bab707f65d595186d"),
		nonce:           decodeHexString(c, "1a63dcfc52e55c58"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "49fec2675885d54a4c6b101f291945c735ad9c2db51a63b941bc182ee51fd8fd84b8c6337b0f77f310ca50693b91b590e3ef655be7ad7621ed21ff39d3292efd"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_12(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "a09091ca3280f7f58376bf680de11192"),
		nonce:           decodeHexString(c, "1d08dc0d060095b3"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "0d59aca023647f0bf26881d9126bab8c7fea922d2b4c24f1dcc0bff87a3d0d1b1da0e875626a56426186efd0071f5a789fbf35fa8bfc85afafd3af6c9c7cd07a"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_13(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "8b4175d0a19539ef3d5d40846cb40ab3"),
		nonce:           decodeHexString(c, "8c310d431cf00b3e"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "92b3b0e57ae4cefc40d6e5fa0d9fa85c970c2dd0cd4e04d7272756634b84cc20c5368f3a7b3e1211c5fa2e6335436b88582d048bc76a7c19bbfec135a1055cbd"),
	})
}

func (s *drbgSuite) TestCTR24_AES128_14(c *C) {
	s.testCTR24_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "68c853b9129cb20a3ddf11e78a5875ba"),
		nonce:           decodeHexString(c, "f5ab548605a51103"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "43cfb03a51d7da40b39468361c2ee078819075872f4f7c5d2b09ef39914602a72a62c63e29383fbb9e450fb2aef32eb9f370cbbc1ab4708a5d2898df8ae4f626"),
	})
}

func (s *drbgSuite) testCTR26_AES128(c *C, data *testData) {
	s.testCTR(c, 16, data)
}

func (s *drbgSuite) TestCTR26_AES128_0(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "aba36ff7a53537454b5cb26839171540"),
		nonce:           decodeHexString(c, "6871c0f526fbcdc7"),
		personalization: decodeHexString(c, "dbc44af498161f1f2af6fce66ccc30a8"),
		expected:        decodeHexString(c, "af686e9aaf10aabcbb44b3748953ad185dbf12898e524d04086102e45f3841c650f623f48f542caa14793e4fcbbcf2e461be1c01ed8f1f48b9704d79a8ebf79d"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_1(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "47afb83e3ada220f7d7e6382a8b38cbb"),
		nonce:           decodeHexString(c, "7ab8ae9bac8b15a5"),
		personalization: decodeHexString(c, "8bec1cb2d180b3677cd1a8604b614dbe"),
		expected:        decodeHexString(c, "90cf1d9e65d976cace2f20e78147d5040d02237e04173f1f3710e5227dcb8564684f2eba38e1def72b93bedb4485f2b817ee66c189024b2a127365bc83500871"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_2(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "bb6993aff4804a87dd425973fde35311"),
		nonce:           decodeHexString(c, "fb14700b335a0146"),
		personalization: decodeHexString(c, "1eb06544caa86e2ac4588aa851e8fc0e"),
		expected:        decodeHexString(c, "4ee727073abaa3d7418d6b3e3bd467c9284854423ba6d4ef02b2dabf14b9b824b27ada2b4a42f7dd1cd39dc4420ee6e843fa7f2eee06bb05c647dcc0d697c009"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_3(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "a05d9d84a50de79701c75959af26a8fa"),
		nonce:           decodeHexString(c, "7ede7f1630dc01f7"),
		personalization: decodeHexString(c, "54f5d6e5d9b1feb0a1c3d9a7ec81dd98"),
		expected:        decodeHexString(c, "f47aab57dfff931d4bfc75a0131186f0f5c1505e9c6c7eb935e31b49f134efc00e45fc967358eec4a921fda0d0537d9e4fb33b263db8e08e73f21fe17505196f"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_4(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "807d1b1096e3d7f789036d644dfb9e8b"),
		nonce:           decodeHexString(c, "8cd77f71ccd8a322"),
		personalization: decodeHexString(c, "4f72745e2f2aa5436189db9227820e46"),
		expected:        decodeHexString(c, "ff4ae0064c25fcef074ab3650eb6d3044f86687e6db22629992b08ed078c65d03daf6241a310e5763c298663848e32d0614e98989a16d37dc1729135fca5e62e"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_5(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "1abce21f123368b26c6656b0cebc03e8"),
		nonce:           decodeHexString(c, "1f7a6d5bc2cb9759"),
		personalization: decodeHexString(c, "00cd259bb87e4e6e213ed44eec19d99d"),
		expected:        decodeHexString(c, "86c00d3181c1b606c51f90989383b4b8922e4a190b9401658f8d4513697ca7d584f5fcceb3324f624c4781dfe55c0dba2d66bd858fb643f23ce5d67e57250007"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_6(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "06d5fb814ca4b2bac9e1cfba0f13699d"),
		nonce:           decodeHexString(c, "0ea87b9f3bb1e629"),
		personalization: decodeHexString(c, "f63679dfa353703f12e7236173c7d320"),
		expected:        decodeHexString(c, "e0c8ff1f1d2c69afcae0b73be8b3c4c7413f7fbcbefc3bcf1e688d2a7d0849fdee60bde91a0fb1a5ef4bf3dfb336b73ed047dfa77477a51c6dee817263a20c37"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_7(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "dd3b7cfc518c180cf289f14525150ca5"),
		nonce:           decodeHexString(c, "3694043c71c2ddc5"),
		personalization: decodeHexString(c, "3b1c08fab6361851d7a518ae355b8c9b"),
		expected:        decodeHexString(c, "59164bac7a714fcd5a4580fb54ed4ddf99c39cf59c23f85f6c5216f4e89cf28da1599f8257a6afc302ed3a1dec003ff450912c2bcd682cd34079cfb3ccf2593a"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_8(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "b7314bacd0b3e9e2214e11a49c4fab54"),
		nonce:           decodeHexString(c, "8cd59a9c39501c0b"),
		personalization: decodeHexString(c, "b6e3a4644fb913a54c8921cbc1737238"),
		expected:        decodeHexString(c, "e92c97ccbbd6013178ee06d01ad2c9eb5464a7e30432b943e0b371f136a994b9f544f37b60561e1025b12b5a15cb661b301b5dd4384b8bc00d1d72b3c618f875"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_9(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "d62b1fc2b045718a3fe59dfac4a1c03b"),
		nonce:           decodeHexString(c, "419917d58a6756a5"),
		personalization: decodeHexString(c, "f4794bdc109f13004b893b3283c0977e"),
		expected:        decodeHexString(c, "0542dfdab88e349163ed9ba634ee76257dc9276661cd5dd2faa931bc3a2e9c2d17e570ffa1a5f14496f0eac3339efde46aa40e87f21a985495fda394f2066ebb"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_10(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "91b4aa7d565832e96a21d680a0ecdb4c"),
		nonce:           decodeHexString(c, "060909e2ce8b2dc8"),
		personalization: decodeHexString(c, "5cbaceb89c150d229b516c349360f27f"),
		expected:        decodeHexString(c, "c079c2917e8f6b84c58e0226ad0b8a60a87b8822a990459de24cc6554c7f241aff30cdea61a748470a58c94a150ebddc355c644dd4786e36f17ecfcea7339d42"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_11(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "32e166c0938d6da721cfce235a66a180"),
		nonce:           decodeHexString(c, "eb687996a8ff6394"),
		personalization: decodeHexString(c, "05b99f8d9f102c5ff6d0d45c4d685ff8"),
		expected:        decodeHexString(c, "ec5ae397257220452e6105a5a29f56d9332132f5338d19e81af6d31022bb91f87f04e4520d50766143d3533e7b9d570a806884a3dfe281e000f77e7b85594434"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_12(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "3a3d6aa9db0f540a37eb81509ddb655d"),
		nonce:           decodeHexString(c, "1fe8c25b27ee3abe"),
		personalization: decodeHexString(c, "8e15a5ad7a5f2938ebdb5c088bdcf307"),
		expected:        decodeHexString(c, "8d7eaee50e87f0bc34a49be94e7a4a64fd1ec9861ef334222ffb15af8929f9ed42fb3fa629d6ba39706b9193e12ba2404476177c4f69bfd18cfe59555dca7c04"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_13(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "c08469f836fc1248ded4edd8af4d9de8"),
		nonce:           decodeHexString(c, "746ea2d74cccc4cd"),
		personalization: decodeHexString(c, "a51818ae5437562552651180f3de5dae"),
		expected:        decodeHexString(c, "a08aeede463fc9ac48907ab4b0b39d6f86a879f421816141a052eb48d87ff5d9b5b6cbfeefdd8cfc1772eb626712453f88ec747f6a05af917c9fa161ede11d7f"),
	})
}

func (s *drbgSuite) TestCTR26_AES128_14(c *C) {
	s.testCTR26_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "0ce4c1828f9c9f0aeab4953672235923"),
		nonce:           decodeHexString(c, "464b42c9fc5b3831"),
		personalization: decodeHexString(c, "48de34d848298dcf3f58c52d96c77bf5"),
		expected:        decodeHexString(c, "a7d3790226f5ab5b3833dcc763c2e9e7c5b77d57c2fa261547ae0e39e9784df269d08bbd4036462f3accc71b378b0941e99c327c4a503439a93b4ec7a039a2a8"),
	})
}

func (s *drbgSuite) testCTR28_AES128(c *C, data *testData) {
	s.testCTR(c, 16, data)
}

func (s *drbgSuite) TestCTR28_AES128_0(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "6bdf5332bdce4655d45c2cfea897b000"),
		nonce:           decodeHexString(c, "e78c5571c5f926f9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "e0715688765a3285e7b7db555f277924e7171f7541bf26122b13dbaaa39f9e2b0345c659583ff8c9cfd888f1abd2f3b36a7c9d47c687b01c819a9f9888542e0f"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_1(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "a4737d48a89325078579e649e2fa65eb"),
		nonce:           decodeHexString(c, "6a799a7a2f13e813"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "752a916d980518d9e7e47599066b45714661f34159f0c8cea8dabd596a066aff7ae6c21e69a356dd2ee0da55429c675aa6fa0900173f5477cd7fd649eae0c99a"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_2(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "99b4e9e5baaa99bf8b8a0c2cfd9a89ce"),
		nonce:           decodeHexString(c, "f48b312941d3554d"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "85cf148c65ece7525dea96344da6451199492185c8b8dfff500ddd68caab50bd741866937a501af876ae849b5b1b5330de65ebf38e9d559693fae805ccc9aaed"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_3(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "7c1ef846e7f3a7cfe5c7bb47e46aa0e8"),
		nonce:           decodeHexString(c, "78eb002955d270d3"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "b739f95339e83fe8a339162f1e96cc82b67af41759f483064e61a039137af07d934a0eb7ca7284c14686ca0488953ee440f6cd1145d0766f4ff1c77b0d81d361"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_4(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "6f750fb73fbcb6db57506cf2defaf6d4"),
		nonce:           decodeHexString(c, "a52c9297e02f4255"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "0ca748fcb3ba23bd043f48507585425234ec4a0f350efdcc87f9062c8ae0cf1e033b7df8357f5b0e4f7c21d4fbdcf89ce0199c25790270fe67ffeec26cfc4d18"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_5(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "44b6c39a3af6a473148e32d9792b9eee"),
		nonce:           decodeHexString(c, "de44dd36003e4822"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "101220087ca770e4574fd05b2e88851b4809cf21e65eb0e3d1ecda29aacf2d93e75d793b9924aedd9b9a38844db430470d015da60418753f2e3c6a15f558f4a9"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_6(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "e4d1ff1c3cd27849a297e5a98ede086a"),
		nonce:           decodeHexString(c, "f89d58b0617e5d10"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "54c46797a38e11c543477809d9d6ed0ae4020285edf99e7ec945cfcf21ddaa6d9a71856bdbd6613bf8e97e0af8d3c4cedd0dfcfb4742a2ef0a443ef878960e6f"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_7(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "1514be706e4c1672905506b0c55347f2"),
		nonce:           decodeHexString(c, "ad8d889f2b91a6dd"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "b3644a6d16340ef8b28ad06f4cbfd5e799f080d4b2f831fd90b313c862d5763d22b1217fe3d4f07c006a39895d46d52470b3f8b77521e141f68b06b54c38666f"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_8(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "8335a433a80d5a3e8ab1f73855f97a9b"),
		nonce:           decodeHexString(c, "0262b47c2ab5fcba"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "66a03a0b3b2ef8d04e01da6e31bc9002ebaca98d63fddeb07fa3e0650a32e142fcab9963f06ba88cd58be75a99dfa13b4f8feae6f84ce2de4d201992231ff3b8"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_9(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "37faeb0b3f85372ccb4422e36690e96c"),
		nonce:           decodeHexString(c, "da165c92d5faaa3a"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "4775159025d5b882f37c7edb8d89b38acb14a02839802c377ac482f8c6b59ad69c7d67556bc41cc322d9cd75032af52715b3602f24c2b099840eab00b74293ad"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_10(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "b6830f64e47388a9c3c64aa0be708451"),
		nonce:           decodeHexString(c, "5cd43034c989da56"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "e12f1e9fd90e4f166a13a4d35f9091a974443291ffae4402cbcd9a24893b77096b0a00db20e43288626139f7c1767f6e7ffee9022dde41b478485ee9125f3eb8"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_11(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "f71a01ef00b3d6134607dfc57325f687"),
		nonce:           decodeHexString(c, "49760aebbc8961fe"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "0f008e60ab4814ccd8d2897753f58ca358bacd3f2efc22c0ec89a65b0a9182f555cc497aa59af914bc9c65be7c092146cb78fce240fc8fe136729ea77716447f"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_12(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "790ef8f0ee6df480b2d44dac1c37614c"),
		nonce:           decodeHexString(c, "d552444d7dc35e5c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "2af8e30977556fae7c18ae12ba2201b40de8c02166ed94c141272050dbea72f3b8f991547fc0f55836171267d42a53dcfb37c15a227676bf218a49d0b723689c"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_13(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "be196c9ce383c8b101cd27a112919fbc"),
		nonce:           decodeHexString(c, "332ebe3d473a3a35"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "016b05d557777b36a192040c580a5c516dda1155934afbd96cc85b6e8b0083bbda273ea6db52a0589d9de2e5568ffef9db3950093c4db2e8ce9c2513084693e4"),
	})
}

func (s *drbgSuite) TestCTR28_AES128_14(c *C) {
	s.testCTR28_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "7639616ee5cb269b5f540de150b4650e"),
		nonce:           decodeHexString(c, "3c933abed307f6e2"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "0633c9af4d40f3c21ad4e078a0c84ee03ad7ad410a0f7cf3dfdff5920dcd3987fceced11a2b38b15535e44b55377a8f20f4fe05187f976a276e64b813e55a8b9"),
	})
}

func (s *drbgSuite) testCTR30_AES128(c *C, data *testData) {
	s.testCTR(c, 16, data)
}

func (s *drbgSuite) TestCTR30_AES128_0(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "f41f466b3219be21597763fa7b76fb40"),
		nonce:           decodeHexString(c, "cd93feb9962e81ac"),
		personalization: decodeHexString(c, "b58f869ad0aa9808f6646137431d430c"),
		expected:        decodeHexString(c, "2fb6d7eca392674fc722a619202e819d0da9d11bc67db10be4c13cb964e30ada96dccf0c922b710ac00ded5457fa971bb1c661a09afa720a5864344bf77a36ae"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_1(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "9d1b8834832ffa13832eb086047bf3b1"),
		nonce:           decodeHexString(c, "d0f15efe86477f75"),
		personalization: decodeHexString(c, "73c93734f6ea39ae04e6a4b49766b820"),
		expected:        decodeHexString(c, "9fb67d35378940a5d76b963ace4f8158e93fe0ca064f9656d46df1c10d025f48b33569da07c77ec512236d08d26997d6b9bb6915df639ea89da957e66fc29003"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_2(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "2989862a79e255195a24828fd300eb34"),
		nonce:           decodeHexString(c, "30ffdacb3ac7b27b"),
		personalization: decodeHexString(c, "719b899c9e4a5db9e71dfda48fa658cd"),
		expected:        decodeHexString(c, "b4f20060ea3001efdbd5cc89838e0a08c09f7a6fe5bc023c33d115fedd6ae151307422f997d32b3ceab87995862368c4c3af7ac4815874c0084ea1dcec5058ba"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_3(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "7482b2b02a7446de589d7e60cf012969"),
		nonce:           decodeHexString(c, "1c4bda6a57f41ffb"),
		personalization: decodeHexString(c, "529e4e80f501e73ec86eaa259628a197"),
		expected:        decodeHexString(c, "7498f9172af7f5f26d847797768e459170dd9ec7f42a1fe979a2e4fa32a5e124c5cb1ad4c394a2c2099e8f942efbe59af0975b56a9afa774331612ad887b3f55"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_4(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "0a8405991aeb64f3a82d8bef2b6c9422"),
		nonce:           decodeHexString(c, "a7a58da9b216f7fb"),
		personalization: decodeHexString(c, "68df62fc01d3dbb018c163be3429f2aa"),
		expected:        decodeHexString(c, "7a9ba1e825133ff4c1d646ce5577f35a1784eec2c1977090b48e30bd3b7506f447ee62d021cae12ad287b417eddb9ec6460e3e284afa73b739564e4073d00e3c"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_5(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "c7e9f68833b18e036aa1da025a359ed7"),
		nonce:           decodeHexString(c, "99f0e49ce811ee7e"),
		personalization: decodeHexString(c, "8de7c86b8ec61d6abb52a7916671adb9"),
		expected:        decodeHexString(c, "1407b68151fcb0f08ebabc21c6c181ac1dbf9c6fb1b2c16eaf1f8c490d6f7d52d0f421116a5998330d8105f5027617dc94b14c083f49d11c34f4f26302316624"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_6(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "d67eeb00a4291a18471184b43159d2bd"),
		nonce:           decodeHexString(c, "e4a361497ee1438a"),
		personalization: decodeHexString(c, "f054adf2ad3849da7272b3b32ae0fcab"),
		expected:        decodeHexString(c, "5245c30a651d9861b636c8e8cc8b845246da10c104d78134a84e41fea80e0e73692f8481cd4d750c79e3876b9abfa9d614d868249605619defef2968fd33141e"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_7(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "ce08f02bcde47f6030820243e1013887"),
		nonce:           decodeHexString(c, "d1b15130cd038d6d"),
		personalization: decodeHexString(c, "ad0bcb8201160d82f17966d4c7b6a4ec"),
		expected:        decodeHexString(c, "f48ed03bdba9310d7fe0a5dddf9df49c0dbe07b95bde25a0b66ed01a9f7a07820f2d7eaa986370a0ce00013b4331e44beb3010575af7d625bed55a592d973828"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_8(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "8dbcdaa78addb285dfdb5e41eef67e6f"),
		nonce:           decodeHexString(c, "49062c098b64e7f8"),
		personalization: decodeHexString(c, "7dc14e7396f62450bebdf2ebf017aad5"),
		expected:        decodeHexString(c, "49d03fe3b72d44a8a1e469d25a145ba64d6169dd947f8793d5ef43de7b2394137083e6e769bbfde9600c36b032b7786522070b5a65c793926892b9fb0d1c1d54"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_9(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "0d98370cffa0a879fe85b5de4e69cf7d"),
		nonce:           decodeHexString(c, "9350a86b7a1bc68a"),
		personalization: decodeHexString(c, "0161b24d1c3ca590117ec29acc39446f"),
		expected:        decodeHexString(c, "7cd64684289b43b1e593d94ae979f7fa5de89577f37c7aea0b584d18602c251b1ec25ff612d97fda3784e9655973e31fd4b5d1eadec66861ae2e9719ad344730"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_10(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "d9b6fb62da0a023235e7cd7816377cf5"),
		nonce:           decodeHexString(c, "1ca8c5fa5b9f8c90"),
		personalization: decodeHexString(c, "17be4885b5f9641bf5f689bc97978858"),
		expected:        decodeHexString(c, "e11e3ca6c832f46ff1f971c07e7b66c7d5b1b2e6ec8c5ffc77103f0ad308800bb989b9ab7010683378a3f7819a297a3765256ac40fb0a9cb2246aeb85d73601b"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_11(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "a7113944fc4d4b15a94967c8f5652a36"),
		nonce:           decodeHexString(c, "a506b79afe6af822"),
		personalization: decodeHexString(c, "918fe4f365259c18eb6850fbee403f5e"),
		expected:        decodeHexString(c, "6076b2c0f7de1dd3e46adb1161b72a7f83b8773fab0dcb1c1fde25db0d09862acdd38f8d2164903a8e8858f8a9b61bd7dae3f60668e6ee264b9ffaf578b55646"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_12(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "acf7623839b94bd7a893c287616cdb6a"),
		nonce:           decodeHexString(c, "a1f6039f91c17623"),
		personalization: decodeHexString(c, "5f9c69de2c32873f679d1768ef3c70fb"),
		expected:        decodeHexString(c, "f6abbffd7965649436b48b092086cdf0502f52c87c8fdc8f6d49f08433a202d8be44f6562ace580e8075f0bc670b2e9d951f15f84d82afe6b832765ecab037ee"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_13(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "d90d786e9d1962843b027aa0e8598fd4"),
		nonce:           decodeHexString(c, "0417f14f96f2029c"),
		personalization: decodeHexString(c, "ec5c55145a4f9552f251f7dfb017c1d4"),
		expected:        decodeHexString(c, "fe2ca44455a5e7de708a710ab646cfcfbc87d678b3e941063fa701dac7cb4f0953af8fce458cee549f537fa771b89b0494c0bb5cfd35bf67f4b270d6205f4703"),
	})
}

func (s *drbgSuite) TestCTR30_AES128_14(c *C) {
	s.testCTR30_AES128(c, &testData{
		entropyInput:    decodeHexString(c, "1cdb744559a82c62fb3b15a2a3aad189"),
		nonce:           decodeHexString(c, "43432ce30ae7f96b"),
		personalization: decodeHexString(c, "f59ad9939fcd6c0f478cc50839f8ffce"),
		expected:        decodeHexString(c, "91b9ebeb929dc699949122970b2177d5d5b705042d3a0d60d7d3bb218b7a69e1cd6864b8a268ca78c834232d0cb88f937730084b1ec7d2867fbc850e04050324"),
	})
}

func (s *drbgSuite) testCTR32_AES192(c *C, data *testData) {
	s.testCTR(c, 24, data)
}

func (s *drbgSuite) TestCTR32_AES192_0(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "c35c2fa2a89d52a11fa32aa96c95b8f1c9a8f9cb245a8b40"),
		nonce:           decodeHexString(c, "f3a6e5a7fbd9d3c68e277ba9ac9bbb00"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "8c2e72abfd9bb8284db79e17a43a3146cd7694e35249fc3383914a7117f41368e6d4f148ff49bf29076b5015c59f457945662e3d3503843f4aa5a3df9a9df10d"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_1(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "c15f9fc5741f2cace0b58d7249bd0377bd5708e365884b59"),
		nonce:           decodeHexString(c, "c3343e3a11b2dc15261c51751f513b60"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5c6ced8050bc5ade3f9acacf23899f98f5e045a768ba538fc371747462eb9b84828c9ce88e4199052359b15833668944d618767d0c6cfc2411f82e0412067af6"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_2(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "34dbd5cfe79caca4258cd1ae24b5afe2d37cb4695d655a7b"),
		nonce:           decodeHexString(c, "55c28fb44e5849b57a77873513c65705"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "9d471d2cb8e42002ce24e0c6be84c89937595e61faf0371cb7e8285d76be6b34ef1db5f7d2ebad4bf8ff77651389efcf13bae72bd8a13c8e9fc61855581b8ea9"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_3(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "643e95bfeacafc7d721b48b9f15a80ce8ad7077f301f16c5"),
		nonce:           decodeHexString(c, "2d908dd6952f862ada902bc1a14a6129"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d6ee5e3796349d47c2ede32c938f2c06ef8a6511b4e9bbbe74eb71bc3c4770d90c70cdcdff6b3c9fbf4fa6bc380ac2f2e16196b4c4c8a6448f040d11d36b4214"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_4(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "0f477a591490a0accf829c673f1506daca70df96582c3d3a"),
		nonce:           decodeHexString(c, "1a8dd4233bede3f2838634c09544ab6e"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "43167573867d53ca92e29686d9fd993c69805a6311e33eb96330c537dcbdeab11b0ba15ee5b2287f2485e2e3a6e0ed62d2eb8108c1fc73908dfd4bc7b71b166f"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_5(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "00a6ed778311528109e6fb5cbd60d1a76c4bd5d00dd103a4"),
		nonce:           decodeHexString(c, "4089f2f0b74c184ac9afc95647cde84b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "b908d93ab2fdeb1a547b6d15eb950eddce25956d8f5ee9a4b1d508c08277dd7a1a1060fcecaf93eb7a472a363e97d7e63cdec7f94a5dc26db7b2b7f12cd8f3ee"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_6(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "5f34552a455a029b4bacf1b80cb1cfb412565fb7734f7e0a"),
		nonce:           decodeHexString(c, "bd55aef06db6f20df271cf60594c22ac"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d4b100ef27d79a5d10031db0ccc72ba89c6819776c19f02f9afe4a4405c584a06621f1bb5ba29af5ea696383a67a9ec652bee75928113db28f39404ff0ba3738"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_7(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "6d7cb232a3203ca6955c1f1eddd6b8e09d820bf6c05ffb6e"),
		nonce:           decodeHexString(c, "c04f1627ed1d7b310aa47ecc7983836a"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "63d6963eff3d50fc0c21aba61d6365742c4bcea7bf76c95d645e88cb139253d563503ec3999ee772c897c9d3663f645ab3041dcf3d9ee8f4c92e895b7644414d"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_8(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "48bfcef7dbe68d5e09df38e97f8977e2a15014bf751f9e27"),
		nonce:           decodeHexString(c, "eea03c6d77802aa2c5d5e1293631e39b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "ecdd8e7c7881946982ff3a11a43bd020e970a3b8901387a1cd3eb63df3fb510660bc82a5e029beee26b4cc5d891fafcf5e42b6f4dc414916fcd8b05ed3f4a685"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_9(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "5be961754fe9e9720635875feb2f37c74fb9620e47b88141"),
		nonce:           decodeHexString(c, "23d35aaceefb412a6fbd187f5787bc53"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "8a8e0e443b8e944d58c8407392296a2c7fc4eb04fcb72cc7ea99513559c0917661db2eda209fbbed400089e3fc48ff6990d97c041bb5752792d130a15257ca44"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_10(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "2a2535a5afc9b1717542928a3d35e95a7f63b877a8388a09"),
		nonce:           decodeHexString(c, "ab217b029a33f0563530d37f4fbff10b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "31a210655529fb46eb4e799710c63192e3656cb1ea240408527df7fef82b305c5fa98a1d824f7830d84a53643cfb52517663d516115d45272f5972240fe84827"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_11(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "ffd8b150b287b3ed9d0128dcc8cc5deb90bda98aca791483"),
		nonce:           decodeHexString(c, "bc7b6a7fb0592fda4abc43eeee42b1f6"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "3fdb4ff83c362b76a83b00a8905f2c058bafe8ebdff8eb2e5f53cb9ca5c82e3217c2c60326b0bba6c91c440a4852c72cd73345c5e0d4cc2e158401045a77e0f9"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_12(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "cb33fc14f2b69401b78bbd781cbfe6585f1e736459000c26"),
		nonce:           decodeHexString(c, "5a914f81d65b197a5d8924fa88fa9ba3"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "6c632091886e24265d123b5c7fca063d70dfea32c7780afaab895c10b51262bc4d363d2165998367c9a8a516476bc763990241baf92d6b87e4de622d4019db6c"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_13(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "c457e7d1c69b1a5300b475ef80441c40800cd3764e414ce1"),
		nonce:           decodeHexString(c, "3933e498a0f20d3fcf05a3b82b67ac6d"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "2689419fa0f9b3a4c5e96a71f676ef7f76767b51b3cd5ce837e9d8286fa837a5a3456695e3e9ca9396e8d5b9478a1e8cdd5c31fbaa5e0743d277fdf4afc7959a"),
	})
}

func (s *drbgSuite) TestCTR32_AES192_14(c *C) {
	s.testCTR32_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "49bc4aa4caba364242df5184909ce588f24a2de340e33082"),
		nonce:           decodeHexString(c, "5101f00ac1899c9ba45e923e07166407"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "9b71e0a09c393eb0f034a12d3ad1320ac1ca422a82967cc54d14ed0e36429ab2a0e8c67a5c7f60e537e410eb5ff075b7c67ac4805200a5ab2fb629b48460cce4"),
	})
}

func (s *drbgSuite) testCTR34_AES192(c *C, data *testData) {
	s.testCTR(c, 24, data)
}

func (s *drbgSuite) TestCTR34_AES192_0(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "eb4553f7141bc10882f0a3741780a42df8bc38a7121d3b7e"),
		nonce:           decodeHexString(c, "6f347f9c1de84fd5341625ae8d6bf50c"),
		personalization: decodeHexString(c, "5e2e73b86ca2f3150d53d23d590acbeedaaf91638bdc3f9d588e945af4bb6ea2"),
		expected:        decodeHexString(c, "a1deb9a5aad608a587d61ce5e0d7c7dd449b8c87898354ad1add6e05853873d279ebe4132fc236a42d8ff0dc3ace95d2cdf9d0b057117cb119ee7550ce03085c"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_1(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f9ce2d0649bc99288af15fdfbc3db88956d96c84c0d7e5d2"),
		nonce:           decodeHexString(c, "8cf00c637a079a98362ead51149e5567"),
		personalization: decodeHexString(c, "b244d68a9b30f3ac88040d6458a625080020535341533be270e894002c07697d"),
		expected:        decodeHexString(c, "2851192fd3b37351d05110974388ece011d10e7b9d380140291048ce3672c134bcb4a0cd074ffff389a02af59c5226be0253e7b7400e6344b1a0d0d145ff366c"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_2(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "618ab7da12a5839b4d808bc27cd5d32500007814a54e5de0"),
		nonce:           decodeHexString(c, "c3efab857f1feb049ee060ba760f17e0"),
		personalization: decodeHexString(c, "33c6af4e264f0d19c361ececf89bd7869fb0af7f9b39159c0fabe0811431e62c"),
		expected:        decodeHexString(c, "1d74ba44eadbae176a0a870622175e4b0ee4e4352f8c2ee19553dcb2100f8e2f132dfd4f4cad5e01e3b702228902dcbee5afd5390939c361882a0b679dc2cd69"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_3(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "7df2a3c0bd95c6ba8873d6a5ad5dbf0e069da3cde1e3943d"),
		nonce:           decodeHexString(c, "fa28224a8949134e2850c52a28576f65"),
		personalization: decodeHexString(c, "e2dea15d8d3aeed87ff45e79a4a760a89683dcf82cfe356467affc44592e2bf5"),
		expected:        decodeHexString(c, "3c48823f4528b396c8667407cb8846a229527d9589f1ceb2fee640a83f93327107c4c991c2c8ba3ee81853140a348c1ca5ce264ef75db4956794268c5538101f"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_4(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "2d8a1a3a6a77e3cd4640a0780f59cb975f81c2733ad7f498"),
		nonce:           decodeHexString(c, "50fc6bcc7ecbdb3d57beab7e28a49bc0"),
		personalization: decodeHexString(c, "95f9c3563b535e69a49134c336cb80fa9ad95108c756ea261f5b3ae9cbafd41b"),
		expected:        decodeHexString(c, "86871f9cb6b4ed252bd1e868c80a263e025bbae2285cca59c629982732a5063e5cbda276f282fdaa90aef8ec6dd31d32b704de5028ddd32e22de3680086f9a89"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_5(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "133524edabd5a7f7858c13f27953e987283cb172a2c37f17"),
		nonce:           decodeHexString(c, "5e6cb7f7c98ee13f467d6bda288b577e"),
		personalization: decodeHexString(c, "67d39160ccee3040db7820a37aefb7d4a10f7dcdf3bca669fe235db63ab24760"),
		expected:        decodeHexString(c, "2d13ee8b73d518b0afbf4e1ea6f632006ba56e3f4e24a21f6877ed0f79c3ec7ffd19ce81bb170b3aa90df697f5cd9872ccb3d1e33894dd16c6f5cf2fb107aedd"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_6(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f297474b5b7e68d5aea948f751bc899b36c212636f288b84"),
		nonce:           decodeHexString(c, "b8a60dd726a03fd0d2bf195d2cb09a03"),
		personalization: decodeHexString(c, "8e32db4366907abb98c0e09c5360c56fdb6f483c84e606f07f5d1d45a09346ca"),
		expected:        decodeHexString(c, "92089094a89fb532f068a8630fa947b8f86eee22d56f22a514f8a6871aa4c808c8c9f47c1354f151c68bc130c8e85fe47e8beac8cb346b8ee2a7e00159a0ea80"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_7(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "70508f6035cdc41af63252af23be67af4a468f45533110c8"),
		nonce:           decodeHexString(c, "77c2bc1a84f4f991796f9c07092ad5c7"),
		personalization: decodeHexString(c, "477382b2f6de44e2f09a1d135a35f1ccb02851f9e1f003f98e022043458f5b66"),
		expected:        decodeHexString(c, "fc634fd4dca7fcd66c38926518a8d6fdeeca07f87e9b9fe5405efd9af9c941cbf156bc5b09a53f5cc8b5bf9437e676905afaee58027bc25ad46c32abbea05c85"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_8(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "2e5d95687b0e9b777f1394f18663e798edbd24cf0c3b9458"),
		nonce:           decodeHexString(c, "02055fa3eb3d1204c253ebf35e3143bd"),
		personalization: decodeHexString(c, "1756635c33f8686b458daaec9b9b4460b8a1d75aa2e300e75557411249abc00f"),
		expected:        decodeHexString(c, "91c7e7183eed6ba66496678b9ee8ec8b86de02fd7c6cb977482f0df4849a72859a804d268668a8f4c94934413a94a2ff0d9d39b5188607cc75a079a7e4846169"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_9(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "2e99703a2bf195a182d221ce79bdbdfca4db53e16e757a5b"),
		nonce:           decodeHexString(c, "4be52fba5d1fbc7ea77b4ae30d16e0e9"),
		personalization: decodeHexString(c, "8af346e52f8c9be68a58f40d50057004b7d75807af922732843696c18290d589"),
		expected:        decodeHexString(c, "23352a39994a03f42506c34ddb8e0d19127ebbfe8239c34a2711cdf1b1beb1cf75402a613c8531d1dcceacdb490073c7a56f0affb3f11f26e8c1901faf3dfb2f"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_10(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "23931f7e10ad6e553a285481acae7cfcb32f644e08c5b009"),
		nonce:           decodeHexString(c, "7bc2b57ede9135b4192c5181f90808c6"),
		personalization: decodeHexString(c, "6ecd84ef10b4e862bc3447867b3e8088e9db4c5be5381dbcb60514051390bfaf"),
		expected:        decodeHexString(c, "a5c115bf8f950b5d18718a9a2269759e78e4fb798937afd06dafc9eac9f5726027d639de08d0b7b008532f70bc48a8f88f637e67ca7fb3fce5196e993f6d3a25"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_11(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "619642e863a8f8d2f7db9c2ee56c13a2f039595c29ed7496"),
		nonce:           decodeHexString(c, "9c50ab67fc7668dec6034e8224ae1d54"),
		personalization: decodeHexString(c, "230b187d67c9312eaf1923526bfbcb319d9d339fe8c862d1a9fa1ea7830929b2"),
		expected:        decodeHexString(c, "0ece188d8d1a37f158a6d065f9073769d750b6d33bf8493df96d99fa98c9900076f7abeb02312ad4e0c6edde99ebd61c396a8375a1d71f6b2086d9021c11a14c"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_12(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "bf4e566b60a592f6c889bb19252b54489a5815f2bd074b41"),
		nonce:           decodeHexString(c, "32a2f8e78e17574d6384e1f539959ada"),
		personalization: decodeHexString(c, "cf758d772963f576b51d327dbaa3517613f3387cb90c85497a9360e04ebfb179"),
		expected:        decodeHexString(c, "f6bf711d5002997a3ecaa36ceb5ee63cae530172890764aeb8afd5d812c368568e4ab0afd2ddf97dc310b1c353e673fdac592ac608e1304f97a5f2578c1ec1d4"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_13(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "18010ffbaeacab06a4c55a38e8e936d79ad49af976b8f748"),
		nonce:           decodeHexString(c, "cf55bc78a1f1ee4abdc554e87b3f114d"),
		personalization: decodeHexString(c, "53a974386ba101ba34ea65b75bc360e9c1cb80c2a6508290a78ae23e5c701537"),
		expected:        decodeHexString(c, "5726de6848ee074f353cf17bf3ee0b09b6b603b1246425ddde1d01b5bf5af0d88888217d59018db3171811da02e667fddf8cb35100369dc9a8cf2aafc5452182"),
	})
}

func (s *drbgSuite) TestCTR34_AES192_14(c *C) {
	s.testCTR34_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f6208773ae06c614107a9865c36eeb2ac2e348ee7cd69dde"),
		nonce:           decodeHexString(c, "ee7c9ec2687703aa103070574d6167d9"),
		personalization: decodeHexString(c, "a85a25e5cd6390beab64ab8b4f534280f3fc7b169eb0a75cb77d9bf06292abe8"),
		expected:        decodeHexString(c, "cd5f73d4febc7fe17351bbb9094729759016bd3eae0dd3d9128697813a0b929dcfce6bfefc1e08deddf617d4e727aaa3c7cb1ff23bf02fcda77d0b1502390349"),
	})
}

func (s *drbgSuite) testCTR36_AES192(c *C, data *testData) {
	s.testCTR(c, 24, data)
}

func (s *drbgSuite) TestCTR36_AES192_0(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "a57cc3b995f35562ba30619ce6c2b51f2217bff014006ef1"),
		nonce:           decodeHexString(c, "e50a312b22d68f320d4bac240d414f47"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "f2cffeb0f4ffbe773dc80472082b3c0a877aaa113dc4d8678b25d8420270e35088cd9eb2cbaad9bc3d5b51865447245a3a78b38c51d0a19ca08c6195587dfd7a"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_1(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "0b610e8979682f44d937c974e73a4c3df95a34b092405fe4"),
		nonce:           decodeHexString(c, "085581c01fb9161584b9f6526f547b44"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "282dfb8a9ed6721229f781a15a7033adbcd49a210a231e17eb70d64ca80375a8ffcb4e9ac6e23273ca9654e671ccb1ae3bb596bf7a8df5dd230bf4a2b39bf96f"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_2(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f93e041748910281a4ff37efe830ac0264e4b12ece0678ee"),
		nonce:           decodeHexString(c, "7d4d6c545de8c8118455595434651bbc"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "a73e59a66de502298a8106939ffaf856af9f172b85f5e99b057d6b4d6ef052f1226700e5e398e6802f2a4afa04e7ff00c46f57f345eb03de8fa7795a060cf340"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_3(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "814886a6b9a3cfc318785b41939f7b075ec6b00fc7ff1c9c"),
		nonce:           decodeHexString(c, "d13313123fbc45ce8b9e8a0a4d82b27a"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "ba2a7c50c6d3e7ef2d849106c6e83310d9ad1c1d0ef53e61c8b955acb4ef4342e1c0f5f22e19bb1eea5d0f03b6d5969291c401fb6c55c99f4252b6ca7f341935"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_4(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "69ce9bc6a266950270bab47f5019f576fb9dee08179b34eb"),
		nonce:           decodeHexString(c, "63025d588f1d08f3811696a63a1002e7"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5e52afd2a86701cb40bd77baf55eff3a6c9f843456ed64d4015332779fe0b26aeac634cb0621c43b3ab71eea2adf96312d385c62c31c0951523260c469b27d88"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_5(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "25469a3c9a37bd9ea91c6f243e444c1b36030152d63800f5"),
		nonce:           decodeHexString(c, "a2b839abcfe4a87368656f43b5bb4e2e"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "32c1fd38da53adbc0a21a828dddcb6809fb01d2830c6a544d12f55fa2e4443f95f005c0ff24c6492b64095e3746f07122f983d238f23dafff7103a7cda57eb0e"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_6(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "b88a610e687c1e06d18ca912866b55d276a789cff9439595"),
		nonce:           decodeHexString(c, "254f4baa2b34e9898d5f78597724ddc7"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "32411f3250bf2a18ced48f522b3cc306afc18e7a36cd64816647a00c221b996a82ee850b7c0a8ca2241b7c6e645b453e01fc35fba54d5baed9228a98a672db67"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_7(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f32bb10c9e0d8f534538940fbaa36844e7f6f0919750e1eb"),
		nonce:           decodeHexString(c, "334b85e031519607e001a287e3258322"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "7558c59684f7a047a1e5929c8a37defc2f1f554a34c47fcd9ddacf88bbf7eac40704aeaa503cff8963295e702435f73c93ff354f529ee2da08c350d4ccb6d45f"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_8(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "cf144ddfc132a6a6e7b5a1efd515c704fc22823e217d32dd"),
		nonce:           decodeHexString(c, "5a0bc0c87cccd908b440d6a4c738fa74"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "865db2d0711dd3fdd5d2a19a65fb931eecd82ba2a7f2e53bd6420a8b1647a41802c04f9b6ba34b85a72b48fde760523e28d87edac9b8ad6cba361f1353fd39ae"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_9(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "ade1908534624c6bc3754de0db75313835a828572cd277bd"),
		nonce:           decodeHexString(c, "05a5b6d21b98050593ab9813268e2dc2"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "efb7b77b76d7bf571a9bb13e39ec46a19e9135ab086d06d59b20f00c78f2aaa848f629327a0d670729a45391f2f023a8f45f1cb10b40c6618317a46afe0ae130"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_10(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "3880c2269082475c16b978c72af1de9d74ec3c7694c519ca"),
		nonce:           decodeHexString(c, "83ead0f2d46d8d1a778690ae512d66bd"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "af4d71c1147552bf8b358e16c9c7fe7f036620cbd61f1432c8c5f270daccb2b9e15c12e9f5921047dc92d40a90e45bf15fef120a3e08421354674c95a63077e3"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_11(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "b1d2c9ab5a904567bd0311c642047de1b39441a5e2d901be"),
		nonce:           decodeHexString(c, "91c8d19bf5d4f9511f6c6653a6585ba6"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "8bc7b790c77b047f7a178641615c295b164b09037c06a8df76bda59b77ed285f808c96a4b49b4594a18299c2a62b1e0d9caaeb0f82c1de35b428038d10ab00c0"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_12(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "19700ce2883eac515875fc91353c9a9d5989b0d29fa2eeab"),
		nonce:           decodeHexString(c, "5656449311e246ee8025cf303d6b6da9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "000217849939ed26c9d5fdd4d003d57c12f52f6e7dae364e73758c7fd1225b17ae66f38379a903e7845446f0998a8745e9aec9b4ba3173ba1b2921eaac313e58"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_13(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "552110486f04b8b173748be09db853023d7d1e6b9e6311f4"),
		nonce:           decodeHexString(c, "d28c45ca09d51104d519d381b00cd712"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "b35b8cde56d5b2a365a755e4d0bdbf7002e4b06199a342f1b98b734311fcab66af1eea2c7fbc2642ac364c920af724e37d45e8c17cd4467b892a45364794a746"),
	})
}

func (s *drbgSuite) TestCTR36_AES192_14(c *C) {
	s.testCTR36_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "b2f605fcf9929b1243c88199143f154266fd8587c90aaab2"),
		nonce:           decodeHexString(c, "781c857f75bb5948c2a22dd757b247be"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "7753797b7b3725c8739ff211a0cc8ae8a1e0f28086ede4d4f836dee57fbd7880f10521242759b9725a08cb649659a2d81f540e93f87e4931d81b59d9cdcf45c0"),
	})
}

func (s *drbgSuite) testCTR38_AES192(c *C, data *testData) {
	s.testCTR(c, 24, data)
}

func (s *drbgSuite) TestCTR38_AES192_0(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "50134a638092b09e78708050dc4044e63c7abd2277be3d71"),
		nonce:           decodeHexString(c, "da8807009d400482c6da8e4e9a4ee33c"),
		personalization: decodeHexString(c, "8e0153aabfd2ca2374bdd1e97ff23ee2e7dbff7e836fa5adb65764f7078df1b8"),
		expected:        decodeHexString(c, "50273aa20febe82685d49a013e75a84bccc7c20128bd098a228c771d08bb5303e715fc30823dae085780d6d28d1071a26508130f3525b3bfd505f07575add874"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_1(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "9b6c7b1c729cebcaf4c4afb9a351d2f6d1b38affc35fdc2b"),
		nonce:           decodeHexString(c, "d046715fee956fa5e1a9d1aaca6ed67e"),
		personalization: decodeHexString(c, "376b8873ca51209c93fbfb158996412bc6cd1f41e9a2093d9ed6fb91ca3ba2c7"),
		expected:        decodeHexString(c, "05927feb60855127b4d6632f1217ef3072a03c03740fdc141e56360efcbbd55c5ef516e3913bb20dc4da9e7998b8b593e3a0215dc032241214a35e5e4dddf1a0"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_2(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "32f79c64904f65169277cf0eaad8cee6a613f92f0863b662"),
		nonce:           decodeHexString(c, "f41907ab3dd3fc0c337581dbf3cf6a61"),
		personalization: decodeHexString(c, "fe9baf7043f15c4ebe9330432dacdac1306d4237e765d5b826360aad3684a235"),
		expected:        decodeHexString(c, "a2ce0afe7fdbeb4ad195c4610cae406883b69cbc8548a82d122c4613a62eb36e986cafce10fc3200aef297a3dff01be3a5df6c8258c0a601d89188d5c065ce1e"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_3(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "b7367a69f91c68c675648370285b0e486a3a970d12c581ac"),
		nonce:           decodeHexString(c, "bd49967a92a0f3ca1731125d335f86d3"),
		personalization: decodeHexString(c, "5e7d45326c6e53f1902e0b5a314e59a04474781a1a4cd3dcf13be178ce737070"),
		expected:        decodeHexString(c, "596f86a78c76d693a66bc07ed1f0d9fd3ba306fda4fb456e7dc6812996e2f7bae45795a90d9a92ab4e060d5b02e1507ac68149ebd6f237d3df2e40a9baf8918c"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_4(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "b614aaf79922f83f67a4d49df5c2405fb5a97126fd792d80"),
		nonce:           decodeHexString(c, "ccc186760d3633568a5e9f38c2db4ca0"),
		personalization: decodeHexString(c, "65347e04d7688d1b0b69a3b181613e6abc7803c64882bf62d410389530062e53"),
		expected:        decodeHexString(c, "9d574ac67cc384a88b5aa15e656fe94bc80bb00bfee7fc79aa2fa7d98e6d8745c0381eab01063b1890edac7ef30e34b5dedfe9b1c7f21d484b2f470cbe7bac9d"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_5(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "a4cfbfc89417d3559886fdaf388d60b78227aead26cd472f"),
		nonce:           decodeHexString(c, "e18cd379480ef268552f748e9cc0f133"),
		personalization: decodeHexString(c, "e6c492784242bc71c582e4bd79ccad4cf6b1124901ea7a4a601e91006786411a"),
		expected:        decodeHexString(c, "05bd505020c3b16dffc511d625bdb36d46ca8a6aa1fae90154469a1c1a2cac50598accb994d0894026e4b383f0f31188a7989f98e9a732953c82a2ff32bf54ed"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_6(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "fa72ab81a63ba4906776848ad1fbc672aec0983c9f09be1f"),
		nonce:           decodeHexString(c, "9d5f58709fd23fc230dbd833ac342d3f"),
		personalization: decodeHexString(c, "a6ce15e6b2ce260f8fb5d3d6b92aa267e1a3236e0910be3a9a06c7ec86de5a58"),
		expected:        decodeHexString(c, "bd9b6f17c2a06a7fa742092f2da81f8249899d13f05b9bc63e3700f81689b1113e1d324e3b3412dcf2f85b469cbce44d259e7153339a4757f87031bf59003047"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_7(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "6f8bd57f523a9a7d1bff9e6e21d934b65e1c00f2b06917eb"),
		nonce:           decodeHexString(c, "74410e8f55fd7d8c6e58c281c9748add"),
		personalization: decodeHexString(c, "4b1663ae4b5758bd5703c9aeedd7c9740ce257f006bc7b68f90f71f63719f78a"),
		expected:        decodeHexString(c, "dc84d5ef8194fa6c897535880bf48476f1e53124b7ad7299cc1ff5e8567ed4d5041ba62a29b6324e4a6940ab8fbaf358e9aa2db45f1c2669b757eb3ee9b9ce70"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_8(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "18a2350350dc88058d9718335526a3925973d3692cec6c2d"),
		nonce:           decodeHexString(c, "8b0c043b89717aca319f96721e45f0eb"),
		personalization: decodeHexString(c, "385f0ddbe8a3d5ba480abb7ad54d2aaea12953df7be1d348fb388386be6651fa"),
		expected:        decodeHexString(c, "bf2745cc69c8e376fe2d60635cd2b7f1324b5829b3d23149eda916d1926161b2988b6d01434889126a5fbecac4bbf39ab352042fa0781a3b0b9d046cb9adbe9f"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_9(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "cf2e38dd52041de7b448632bbc11a99479887d44f3ccfcaf"),
		nonce:           decodeHexString(c, "19c4296480caa9e1602e1c5e8f2136b2"),
		personalization: decodeHexString(c, "64d38be36626e85fda026a2add7f981b2e81449ebaf45bfbb580ab076265b161"),
		expected:        decodeHexString(c, "687758f95c73928677235d46994e902313d4246101695359cd81bd035fdad3e8d8dc91d907630eeefcfa44445b00d1407149a1606edd7284afcab5c869d762fb"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_10(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "75b3f14996b08a9cb21111c10564f1242fbc7ec2c11456dd"),
		nonce:           decodeHexString(c, "87ac274bad1f347e053f5f242e5b2a49"),
		personalization: decodeHexString(c, "1d4b88abffd480bedf4e4fcbd86e2cd38c183c93f2efe6218e949763eafba981"),
		expected:        decodeHexString(c, "8958443263f94b019634e37f7e5dec3b9fa7ce24d0bff61c5b8ee54d0f2991fd2af72ef93b822c55bdd2d20ed3a78905a8c601b4add98e6659b2174c458c2587"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_11(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "820ccc71e2472711e63d4458cc0b8466ba0a79c2439a8b9b"),
		nonce:           decodeHexString(c, "98ed21c67c7c8d19bc21837058306bc9"),
		personalization: decodeHexString(c, "a3dff1630c7fee3d696a35fcb6754b63bdf16c0a849a540559eaf350a8a03a80"),
		expected:        decodeHexString(c, "6701d42fd9be4a6ef9750a5ed6817fd16c06e3791f4e7ed6bf7e0784239eb55bbd5bf8bf757f739a53d9e2cecc85db1c35727ef098d19c09000a7fb63a836c68"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_12(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "2ab22f44ac4193b1111552d295e241edf37d6f4c28bbb3fc"),
		nonce:           decodeHexString(c, "83b2dda79e88acb5b9bff1bbf66f3878"),
		personalization: decodeHexString(c, "a6b0a73204bc0dcc83fa2f480db371eb5b183daf5996edeaaeb09b821a516620"),
		expected:        decodeHexString(c, "690e7f00c557f71429bc36425ed7d6215d0fc18129cc0f935ffb2bc4fb2c2b883d6d95ddcc071fff5ad1996609680a7cdcf5099c88371b6ac06508d352cb9105"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_13(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "e5995b25b47412c2cb2ecd2a32b406124d90ba5e26e65b87"),
		nonce:           decodeHexString(c, "12643ea9fd5740a7386f5f6a01e6fe44"),
		personalization: decodeHexString(c, "3cef845315cd98322e2a28a44468cc14ab680503a6fa89ad4f876b0e2c4f3ca7"),
		expected:        decodeHexString(c, "6129eda9ef3b9eb9f38fd8e9ca93e7319b6fded363ee155c979919c34b98ce399b0b4f99f2836f98ea84d52603f5790da52029a4e7cf8717db4a6222500d7d2a"),
	})
}

func (s *drbgSuite) TestCTR38_AES192_14(c *C) {
	s.testCTR38_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f81b4752d9b34fea628278beff929d179d0444a81b0e1b75"),
		nonce:           decodeHexString(c, "ebe6e2743f5fd9c5b1bceb12f3f79c2b"),
		personalization: decodeHexString(c, "d86a1d628bb83d738614a93c44c08d3717f803a27bee61be24b8edf52efcd368"),
		expected:        decodeHexString(c, "acd301c292bb6b801401ff9e3b6556fa1ca8061b99b8bd545b163318bb880228eeb4ee3ad56b2134337792862da251c5ac43c2a2a5d8cb38c18436839fbab4dc"),
	})
}

func (s *drbgSuite) testCTR40_AES192(c *C, data *testData) {
	s.testCTR(c, 24, data)
}

func (s *drbgSuite) TestCTR40_AES192_0(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "4d5be91999c1dc6ae2e5f6deb563b125bc8439e85f2576fc"),
		nonce:           decodeHexString(c, "11e1c2d4d4f73b9c457fcaa06f4af22e"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "53c2c0237edf9425842ab91c63723616e98bb6c12d16bf8087772d0a080d289f8b4f35dfdaef5b11cd588814c6bc01dc7e23b9bdd39c0aee7407f71054a7c9aa"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_1(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "24d8c60f62908b4474b6a01ec88c995b357f82e20d21c8f5"),
		nonce:           decodeHexString(c, "45cb534e0ea5c1b1a75e6a66d990c715"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "86a1deb9a32863f713c3e6aea81504c7ec766d0b7e5aa800ef0b449cc33408c7f87e712cfd58842c905f8a262a4e2af745fa584b370b17e2af89fcb73d399cf3"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_2(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "3354a953f8e6cc3b64a89fa74bcb45971f14c4d1f79cf94a"),
		nonce:           decodeHexString(c, "d2a8307d9659fd46a073eb33f3481ab7"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "f6423332c2b0d7b6edbd09e590207bfbeca62b4e8d80e819d6773411b361ac63c58d587d6ebbae5e0a211d1ff8de7f901df1533738256e44edaa2bb7c1a14e0e"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_3(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "931837f96b3be2587ce637e332a7c4441737f9bf9b4a5a33"),
		nonce:           decodeHexString(c, "dbab3e4eecb453a16cea08d6fafc2f36"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "a8464a65fd942c2beea20295cd25560121080fb69cdcb4559aded0c670c82838f6ba46b00f5d5935cc77a14a8c4eda02ae198ea87c9c98204d38a66a8c686f67"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_4(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "550182725f545e7cd4fd7b977f587451117cc731925ba27c"),
		nonce:           decodeHexString(c, "de0fc3bf40b61edcc31e0ce6f2011cda"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "ea0dfb57cf5019c73f58618fc0003140a36ed687ac1451a5bde09050aa1b2f6fd1cd337db438ca795984ca6a2b935617d9846699d99b6978fe878c12c3c8f9c9"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_5(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "39c7cd1d9862133a1ad4d45592116c0e09896fa668adc78a"),
		nonce:           decodeHexString(c, "d087829252bf38393cd00f87f74df4ca"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "f49021cd8570718583ed21ec55468b4c45f87a7b8ff6dd5172d937d6e9d9b4086a759aad927ef9e8b7c2659b0f06c2a4f8e35ed8b671470d9dc61b031d7d1782"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_6(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "88fa538aebf8349074eeb45d9567a900c1c97643b79a2c16"),
		nonce:           decodeHexString(c, "bb416691b86362fc9677e3fe50d157a9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "9ef31cf931f21d89868eaa245ff048146059684e2fa5619a440f1d8a195500355fe394afe3d3f9ec45d206b37fd8fb475bdb1ac8800ce54f04c99403eb76edee"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_7(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "81a8b2192d59a25d5480465e508c11931a1ae5a427f6d117"),
		nonce:           decodeHexString(c, "1b610df5af73b70e85d97c159a13bc71"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "cb0bab82b91a727ca3e104b9eacae08605e06d18eea0ee2c71ebabde5c98ccca7edb4793836a65f2838b8283cf4adf49fa1f38368776e5ba4ef7c51437888fdf"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_8(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "a49585ca40ea9800208abef703a4c20bd3c4ca596525682a"),
		nonce:           decodeHexString(c, "11f7e4630dc96e2e4a356f8effe8eef6"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "633a6728bf6216ea303f3b27e4faa52d2da4768e6e7ced7dcd33d4df25ad03e91f5ae9e9e1f19c4dad74e3a1ede15284afccf5148f188b95a0567670e206eca8"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_9(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "bde53ca3e0b0116b611915a3bea801418014734088639cdd"),
		nonce:           decodeHexString(c, "71b4993cafd868ac209db722cb757fbc"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "6cf1dc5214af97c1c6d43c31ace40cedbf3c32bc7f8d6928bcb613b6a9b37011d40a16e8739b2fbafdfe7747c8d5163dc4223ce3303532c719abb169e69d3db9"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_10(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "1240c07f0c42da183ebf8f4413ab316a88246166aa7bb759"),
		nonce:           decodeHexString(c, "6e529d4e92477579d07fb9d2de054d8b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "dcdfa1a5733e5bec22a24c01b6d46ae10209351240dbb42686bca8117daaa2287326942021f198ec6bd2fd634fe8ccd24a562c48297630f9f718623bca0a64d7"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_11(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "72ed4b06e75b59c6fbd17bc726c2e92a220162bc6c1ffeb5"),
		nonce:           decodeHexString(c, "466eeb265fa147ea1f773f98087fbb83"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "46c356174b28ae3a903071643abced784b3168bb20a78a082abc3c8d76ab53673cec2181d40a4294cffcc8a2fa43b89bcb962f5ef3300e7df500fb0c96be858a"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_12(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "c447b5eec41e14a9ce9bbd8593e8a7c0f5559f4b0e117688"),
		nonce:           decodeHexString(c, "4655d3c78caa1c1cfe6d6e1166468b28"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "68d41fe28bdb7de4a761cb6004ecedfa310946075092111e806beed634ff540c38987f16a89be52d35e30759030cbca240bbe9a5dbbff19d150be3a2296b8ceb"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_13(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "2404f9870945bd404fe184b119a8a30371aefa8015d455cc"),
		nonce:           decodeHexString(c, "e5fd16641728b4d46e3f0a43ae348440"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "ea4bc0dec0f1a61cf2ac890d0c7d8db4293291486d65f3ed8b5f758eeee9d5deed2a1d0506febe20bfd18b680c5a8d777c6b3bbb8ca1bffa0f105ee7c1507359"),
	})
}

func (s *drbgSuite) TestCTR40_AES192_14(c *C) {
	s.testCTR40_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "4570ff9ab92e5b694b8aceb3d765f4872f2ea6ad7707c510"),
		nonce:           decodeHexString(c, "a3461bc17d5db692f8cdd8cf791862b5"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "b7761d9300a2218b283b55ce29b43c4e7a63fcc9b1d0c1686efd8d5f0e4031c2dc745f58a50b35be81fcd5e97d70f896b17d3a642a30ba56725d0adef142efd6"),
	})
}

func (s *drbgSuite) testCTR42_AES192(c *C, data *testData) {
	s.testCTR(c, 24, data)
}

func (s *drbgSuite) TestCTR42_AES192_0(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "a79943342377ca018559d0886d43dfe018d630590db1023b"),
		nonce:           decodeHexString(c, "1b96f22bd66179f593809db90fdfa614"),
		personalization: decodeHexString(c, "5bedfc44b372884452367229f1f67e93bc447f8fbee044c31e10967ef0120c6d"),
		expected:        decodeHexString(c, "49ebb36afd563e0772ed7dc8ecbbeeba47ccbabfd0beebb5c99c9771e2df9a33e19c4bb716449eb5e9b66739bdca2ee8ca217ecc1c4ff6c034cdfcadb8f1c7da"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_1(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "c37bcfc799b1a751ca6fc1fdec292ae8cd3a59b615828c89"),
		nonce:           decodeHexString(c, "26eb7bb685e4e912e474115b19fab4a2"),
		personalization: decodeHexString(c, "027ca35a2ad52c9eb87307ac2b4dd0459ab5fef874cd25342752888c1dba6321"),
		expected:        decodeHexString(c, "0db0ad15bee89b26ff32ad9b3e5eea25b3026a0e76e85d576a17400e00e38b82883dd01d43038ce6b5645de58f3d08a238b858978de40906a88908925a91820d"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_2(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "5fd9fde95f759685da08af67d9d84aa0d332235a0d3245e7"),
		nonce:           decodeHexString(c, "1d30eb84edd4487ca3be0f08f81c8afc"),
		personalization: decodeHexString(c, "2ae91ad3bbae3f5f80e180b7e94ae501901c015fb12f50d8ac8f2ac2150aa2b6"),
		expected:        decodeHexString(c, "2aa5b7ef58745946eccfaf9ce1e28c036ad4e99ed5051b51c57466eab0ad0831cfa014426c2c3d4be70f0ec77e0243409dacb87ed340af0b8e39456120640eb9"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_3(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "b06bb14cce0db3f41af3e77910361e7785427c8302ddbdf4"),
		nonce:           decodeHexString(c, "f39d39c305d1727f7dcb6294cafc3045"),
		personalization: decodeHexString(c, "0676b68f21a6c2895211ce127a5748a0050eebec08044ee2fe2693be05e223e8"),
		expected:        decodeHexString(c, "bde9735f36b8d8afe8ac76a971ec9466065d22e1885e75db61ca116628106beb15dde7e5e1c486e4cc66056ee1570cd5349b6c628094c9506752de9f3da95f8a"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_4(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "c2109a5344628ffa5c43460e9e385f4ea4c07aa401ec981e"),
		nonce:           decodeHexString(c, "d242bbb3ba282b9bb0d3648aee33f699"),
		personalization: decodeHexString(c, "00333cfe95af5ecb99f9ab6986762ad2d2f7702aa5ec5cefc43177eeb3b33ddb"),
		expected:        decodeHexString(c, "f8c4ed2c6d7f323828c5e9e50c8631e135d1ca04173eaf15036af8250ef09bdd595d03a776ae066da462c631fc6a66eab375fb0e2ede9a3b5083980f4ff810c9"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_5(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "b7a90d58f3650814ad99d308b23c79d93657c73d2ded5412"),
		nonce:           decodeHexString(c, "8ae96938dffcccefc8d5f934753fd670"),
		personalization: decodeHexString(c, "e319f45111cea35290b077302bd3b5d45b0f7c5e9e1d9729dd2fccc758855c3a"),
		expected:        decodeHexString(c, "7bb131479e6064c7767e17af46d528e4b659c7fabf906c0a7679dd8afdd4aaf104f8e7caab9153652a30e38f80dae10e67caf302bf117104949649c7b20f5d38"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_6(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "7618e5546a5981cb4c72a517f29a5fe8886d4d3e9ad9fcb0"),
		nonce:           decodeHexString(c, "d475704c10c4a1866bcaa412e6bfbe7d"),
		personalization: decodeHexString(c, "ae8c3cde39ed90da31b7320cf35969fde8f82894e222ec2e8aa15d67b55efd89"),
		expected:        decodeHexString(c, "753048b2879abd1a0cf2426f1e1b041fcb6ba3f0710ee8ec410d792fb32bd66b1bac95af81a30901f7bdd42e5f8fe6cd28623fdee9c589bc47de1c6a290bec62"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_7(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "1d38d2430e12932da3c38f44930c1a2ff4a62eafaf8c856f"),
		nonce:           decodeHexString(c, "1885f3bd339f2c96595c3b6d998e5083"),
		personalization: decodeHexString(c, "430c7871901e1656e08fb35eb4ec43984937a75af596c94586a8f02af12b7fc3"),
		expected:        decodeHexString(c, "9ec5998142379d14896451024fec60853418a3ef59d3a6bd3ff1592826b11372f321a98c54d1bf8225bb64ec2f2f4f46ed642aeaf28eb20bc993e6d684fa33de"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_8(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "abb6b76806813c2273aa8ed476af0074a7a473ea548cc1ef"),
		nonce:           decodeHexString(c, "d99177343ea3837f92d9c7367a70433e"),
		personalization: decodeHexString(c, "ef6b87661f2a678ad2d4dc85ba0845ec8eeafa2b82b8aeda4fc61938a8dd6077"),
		expected:        decodeHexString(c, "811ac88d4c9615b535b0f72a9f49107baca239a13f26b97ba733e4af3a24db8a658275ee2416c02f5968c22d13aa6255deb098145105226f4aa80ae3df8c1bfd"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_9(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "defc7c59d78e50c41232f7bebfc5e9408cea266fc1e3c5c8"),
		nonce:           decodeHexString(c, "9ffa0df69793bb6d9e7d78ebaf46f6a7"),
		personalization: decodeHexString(c, "98b3fab2ffe547ec1a418ed98ae7ac669c0701178be655a6faf9962d80aee951"),
		expected:        decodeHexString(c, "a9e6820bf83a8b294146fce2d194019a9959cc584729d67dc71742391060ec6b24d685e545807b1ca9d8a28b68e76af256d435f6f831636c332a7f13807b1f35"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_10(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "6cfaafb31b1d6a6b9bed1de9a993b70fa14dc61fa17ab1d0"),
		nonce:           decodeHexString(c, "add461048ff69753f887ab87b74c25ed"),
		personalization: decodeHexString(c, "e2dedfe6b70fd426348b9838b63f01d5211820bfba78a219a88b9b67adbfa452"),
		expected:        decodeHexString(c, "c60fcf4d362c923bbb70da806c92651b239e32752bf437669a39fa5e24d8a304668bac92ba128a8b7e02df00ba178f76de44b656f3451f21250de778e0405f14"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_11(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "9009148dfc32c4cd2e6709c54a93f57aa75a930a395bf201"),
		nonce:           decodeHexString(c, "b28dfc3c6b339eb8bbaf29e5f8198b4e"),
		personalization: decodeHexString(c, "ebb2e330a05d1eaafc5116e9376220732ccf091e1b0e4017faa94346a83945c6"),
		expected:        decodeHexString(c, "43c31ac49de9f8e774dfc8e437e40a87e42529e625bdbd10bf77bc0a8174fb9dd08f6edd308c6eeba3175bd3496da11436de612a6d47f5d9e26339ebecda649d"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_12(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f3f4e1abac5587f2a8409b1a227f9c92416f6567049501db"),
		nonce:           decodeHexString(c, "c1979637e6b4d04bc970d37e11cdebb5"),
		personalization: decodeHexString(c, "3d491631d74d3d908ac58cc47481b5e0964bb21922061cfebd833820e912a00b"),
		expected:        decodeHexString(c, "90eddd439eed53799ea574c9238473fff030381e1a48545bc07bb2cc9bb2d0eaf23a23f009881d1c2464e6b39f02305b3214d89932f8c237964982c9a7adcbfb"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_13(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "dc2465257d558989a340263faa86ef3f98e8a9fb539804a9"),
		nonce:           decodeHexString(c, "1ca1608ef5f26bdfbbb1b506fd772426"),
		personalization: decodeHexString(c, "db9ba18286429465f73b35739e6c987e07c0397dd603e036853c546e4993fe89"),
		expected:        decodeHexString(c, "fd35ffc43051435712374ad903a4655b36d3754d78551c3f48d70494744ee2f336bd19b780c8cb6e5de4a667117d1fb6b81f03d3620bb5be1802152842c6973e"),
	})
}

func (s *drbgSuite) TestCTR42_AES192_14(c *C) {
	s.testCTR42_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "1bebd5d876e0949d1d8deb410d3a779ced95e70c0b4ad531"),
		nonce:           decodeHexString(c, "4c1db08871719623eab2a24fcb8fca1c"),
		personalization: decodeHexString(c, "f73f41c6146f2c3d6ade95462a70f123b9604bea02fdca6b1154c1f86b899b48"),
		expected:        decodeHexString(c, "4a0e35cf6a1e199e56fc6b5f3784b28c1ac70e91fd74d8a7ef45a168b3aad28f324411ee886bb739bb814d5ba1dd794a0d372eed5f4a971a0c5542b947c4e977"),
	})
}

func (s *drbgSuite) testCTR44_AES192(c *C, data *testData) {
	s.testCTR(c, 24, data)
}

func (s *drbgSuite) TestCTR44_AES192_0(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "ab8543818545573ca217ce4b844b9c3966703620784f1eec"),
		nonce:           decodeHexString(c, "4d1ab71f1824560af0deb865ba4b6620"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "cdd9abe9526bc9180cf64ba2679d4c101a5a8b5244f9322aff8a92ed1d48a77ce20e39d1915e9a5275e8a1fe7a5aa8a28b0642daae9a70dc9ee4ea76ac038274"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_1(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "20b01e9604d26326b86c4bb22b6c8b974e2a42f5cb9204ef"),
		nonce:           decodeHexString(c, "9f9d96d1250107694565f50ef05ee2d9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "8c50e736583894891032e5c3a4f509854463687cd1a4d10b77760bbbeac83bc7d9b600aa2fd3b1c24210ba25e216ec4019aa7f75b74d506ba0913faaabd011bb"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_2(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "d3ed0d0bb500f735ba0896c51f7234d57f3298ab7271fb8c"),
		nonce:           decodeHexString(c, "336dcd6bfb58ea093b923eebc844ec64"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "e7d9c798115ec42e0af03e91d89d663231dfb8db1187e43717899e1f809f86fc0d5985709821edb8f60545bf922a828e8ac9c5ba3623c2cf81671ec33677bfdd"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_3(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "fbf3c267264f54e456cc56d1e0ae7fd2e5847499716e1580"),
		nonce:           decodeHexString(c, "bf60b71cb13ffbf28d20f968230c1711"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "ec04ec06aae20e31101e8f3e8057813b15e048c6bf05332c1082024bde43bc69a3f9444f4ae98efbabf3d986327af0c93b17f5d4e13af7a7e219bc93b6d259c4"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_4(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "3e191ff892644d77031b24bf3dd6dd704a9740d2558cb9be"),
		nonce:           decodeHexString(c, "88891fb25e11b86aba15c24608fc0ecf"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "967123e13a69b8dab52f087bf4a21e438120d1cf0e8efc7c1220732599ff6785e02e0df6df95c50fb7cf6cf57f13adae64c2f8ac00f538eb4a7d24025e0501cb"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_5(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f3d7d01a2cb2731bb1d6e99a16745f319a617ef95dcb550f"),
		nonce:           decodeHexString(c, "8a4d890e3230ff9a1eea9a66479e926f"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "3ceef60f8b23e50e7d49a30ac2c4fd4dcce3a16f94330ae7fff7f27f3d569f25300d59fe87ff7d75ce90cd99e94ba4f4ca30d70c759793bbcf7bce44c21dd842"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_6(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "d7690e2ace3c174e9d1c8cdc1c9cda65831526611105198b"),
		nonce:           decodeHexString(c, "b0a11cfb2744ae7b0d675ed0bf2634a9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d3ea3dc6296e3640f22e75ee27e9f7f3c7ff06f1794d915e4fe9a471431317d09d80171b1e7fd8fc57d8b8d1d1fa61162bbfc56743835742b4d526666cf5cc03"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_7(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "8251bb488097e7d07f41a768c1f37c421180c98ef59b730b"),
		nonce:           decodeHexString(c, "b72bc52b358c7933a316b17cbf07cc02"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "31f357abf5a2a021ee7e0efbad1fcb58118e1bebe42ef65ac8cc93d6d0e54f5af7c4022f343ad63efc94d950354b838a378cc77f20f759705bce43d8734e717b"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_8(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "abdb2ed2e09aa00d4c394f80fc0d5bf213b4283b768b5812"),
		nonce:           decodeHexString(c, "89e19db42a962a1d35d49cb8c643b713"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "067ae3ea2304917147cff976864d9b894fb38b1c8b0e571256474a8e3c6020996d169c0415454664976adf9505c53b5e946b4d0e65066750cf0de141e8460b0b"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_9(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "4a44101f7a1f22b17e7ccc91e9789fcc6f1b4dc30972567b"),
		nonce:           decodeHexString(c, "cb9fc3d452f6cbe5c9831e6537027d94"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d93b8586d61091e3bbef187e5fed6a2b1700e93410866d10bc02d3a622a0a8b38d8b083361ad53197bc1811768206e541115bf96121965c16d32e1c1780e4f24"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_10(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f5a4662d9f5156d3bc282a4bf82e5d97731d36c00179adab"),
		nonce:           decodeHexString(c, "3e5e203862bc328e9987a721897d47cd"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5e08ab01b0a4ca992da996f164341be4e201ef0d1cd049eff660595a70cef1723179cc588ca52de7efde206a15ed7ab414ea7cfcf671a05ce953437639a58c1d"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_11(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "10b6e2733ba85603dfd5d5aa5a206752a0f607f9d4d3d73c"),
		nonce:           decodeHexString(c, "c0c280ae014df200d02ccd5b79fd81b0"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "9af46dc01aa860d9f1ea68d3ef573317503e54656363b570ae263b3760dc174943e1815f972cabbb42e600901929553f76739e2d29f77c5dd113138cdf97113e"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_12(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "55fb21939ebd06d2932b6bce9e023e623555f27d73fad83a"),
		nonce:           decodeHexString(c, "43659f6bbb52f416ae9417908beb7b2c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "41e80b17f852dbf0e31ebe5127f8c528d067449ae6aa03c6825bb4c5dcc6ad5a728412133077c68cba1e4033d719856bbb30d04b82d840fb5a91057c43da8da7"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_13(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "93bf67eda37bf74ba84b435749e5c3124b989b6ac6ad7261"),
		nonce:           decodeHexString(c, "294cb26c39586e4895ba8ae779105dd2"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "c4601178924d243fee92557ea30b9d48b816259642da4cbe4d868a948ca7bdbd41e6cd5eb800bf44ed606d2ddccb34dfcf3baca706061f12621c799f551f448e"),
	})
}

func (s *drbgSuite) TestCTR44_AES192_14(c *C) {
	s.testCTR44_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "c8036e05b5c97e1f170ddd6c958cfe188d4474b11d678262"),
		nonce:           decodeHexString(c, "a1829e135031eef00a27a6fe02248c1b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "15be79f77f5c184310416e92d7a47fc32567e0c4858540bc540c819504b39bd82bcb961eaf50e38db90c59309051717c7674fa3e75a4f5ec33669c736ec43a19"),
	})
}

func (s *drbgSuite) testCTR46_AES192(c *C, data *testData) {
	s.testCTR(c, 24, data)
}

func (s *drbgSuite) TestCTR46_AES192_0(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "2a9b56c35d17a5ebfc5b62ae44e929ac3a0747907c15efa6"),
		nonce:           decodeHexString(c, "8df8ca01196719e526ff2ffee201ef45"),
		personalization: decodeHexString(c, "a4f5fabed064693913880e33f5aec5ed132f429fdfeb226b0e834e72d3ffb449"),
		expected:        decodeHexString(c, "51e19a13b11815b1ecb065d54bbfa45e31d94adeca33856254f43481944513de8fa6cf23c4fb24b332346d00464b06e9ae80d98da9c6fd3839cf0ca7531ccb89"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_1(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "8608bdf7d33d89c09324b498954110b8c0eccb520cf86068"),
		nonce:           decodeHexString(c, "572e5816ca90e029102d5d682189d856"),
		personalization: decodeHexString(c, "c4ef4c15721337209ab0c103dbeedb46329358afc4af0ab74a27820088cebb5b"),
		expected:        decodeHexString(c, "c393ebef36f86f0faa9e4092ea0aad0d8b81920d762966f31f1ffdc90fd21306f9047422a9de67eb2c51ecfb27cbd068648596c586d9c94f31e9e82047cdaafd"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_2(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "41c69833968ced2e0d2685e068ea7d59f6e673af2b7d3599"),
		nonce:           decodeHexString(c, "c630e48ee0d5353fc7ed4d607def8131"),
		personalization: decodeHexString(c, "49cc8310e0e89846b64f58fbe8e2bbf68cf999126f4434df3e5326f6196efd41"),
		expected:        decodeHexString(c, "f7f4bd0db984a238dfa9dde2b519ddfd0eb99a54fb0abb73873b3c1f5084dc6412be353f580be01cba58dc1d90a58bc963bdb77d2170159987259b60c66b8610"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_3(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "c1c7216785a91e2816c8d4d141772d122ac3a108d6ea0344"),
		nonce:           decodeHexString(c, "8e228b8fd6e4671609a246082caa3f2b"),
		personalization: decodeHexString(c, "d40f2072a27354dce5d4b190c03c79c70441f1b9ac0e61345b7671dd577ff4ab"),
		expected:        decodeHexString(c, "55d462eb21083b55fa5ba9b606fbf6669190235cf112123a40a4a2cbf24aadb61814fe50358eaa1377fc829312d9834dc6a3735d936ab30557532b69cd7b268e"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_4(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "b3474330a26361abaf69511f920180692a10b829e60f9e27"),
		nonce:           decodeHexString(c, "be420ff26f33ae5c1135057e379ce8f7"),
		personalization: decodeHexString(c, "2b21b848b8aea51b4324f9bdd62752f548fcdc5ab320cb33db1f23f9383b5765"),
		expected:        decodeHexString(c, "314015588ba18c375febd64ee95423d42a7679da10e7e117d7b97089bdd20027010d7e602d6e79b8ec41a7e2bf6ac46a41ea661601bbf666aaf99fb5467f221e"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_5(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "a95c12ce08f587fdecc69e8f41e7b71790094d46dd5ba7d9"),
		nonce:           decodeHexString(c, "be3eca668b3eb25ddc3336ac70e34a56"),
		personalization: decodeHexString(c, "01bfca956d0370b8509ebd4083487228ea9572eb52c64e83462c090772088760"),
		expected:        decodeHexString(c, "a994a0b97c21cc3551cdc64602da0b1f675d7229f450403e303d1a7cb9376e9b5f88992c4386ab75864c24c5bae2fce8b00e9ba6a62f176209894565044c9421"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_6(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "fdbf44186451e19c0e0e71984955d301434614679ca05310"),
		nonce:           decodeHexString(c, "c446f2c50f8148b4435fa7f03901ad16"),
		personalization: decodeHexString(c, "0c4015317c922fbce0934bc296ec4d0c5e87940cc35b535ebb3253c5d10cb50d"),
		expected:        decodeHexString(c, "0d20843219433b903e8a8e70e455cf574bb315308cf9f2d837b7b235a95f2ac989661cb865e8f4b77494a3dca3148015abde0a73c656a7a36aae91f02dce30bd"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_7(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "3ef94fda2d12a076ee5abfd7ef39e6c7d75a7fb9281b4b55"),
		nonce:           decodeHexString(c, "d81072c46efbd91e9c6bda13ccc87c7e"),
		personalization: decodeHexString(c, "5711a9705d800fa1a21bd7752e3f1d3da90a1f3369fea403a9cb6581244c7de9"),
		expected:        decodeHexString(c, "7d70722a76825264b37019b4da286c7347890f59f0a498728e6359bd9d04ed96a7588166c6ef391da07994d63db8a2fb2cdce3187b56aee9a2d5cc0f852b2ed9"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_8(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "f3519e7da5797017efa1a1a015f9a61447fa20d0f2fa4d51"),
		nonce:           decodeHexString(c, "fd2d1eaee6d8455a03e83a0f80e15b17"),
		personalization: decodeHexString(c, "1ab108ec099270049326d1dedeee526b29d17773f766571866f4e9383c25ccfa"),
		expected:        decodeHexString(c, "d2e76643e8f9ec5a09f86655fc65738dd0d20a525f1497b32729afcee873c8099f6b61c4d9324e02a2d89e50f37584c2cd2776342210da42058f07df3f85f8ac"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_9(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "41072a1c20920ce5c52c9eb68b87f97a658a6697d4e967db"),
		nonce:           decodeHexString(c, "cb03efaeca21534c59a9bb12f146ee21"),
		personalization: decodeHexString(c, "9a09ae50f76149a9a1f55a37ec30aa08cb573057d4f5b5fa7713ef7a1c089baa"),
		expected:        decodeHexString(c, "527442041c230ad49682d7892c8516f658ed0b1638d1ed01b279147a6b4a815e4166679099fc2ea1fc44b4e5e4b6106f6691505dd062ea3b3dd67332ec482ef3"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_10(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "fb67de62271adaff78dd3c06fab6b9bd01a8496fc9beb6a1"),
		nonce:           decodeHexString(c, "1c5f50e7372f8f8d4044674c1ff19898"),
		personalization: decodeHexString(c, "8e4f33ea52a9944383ebe90a3042aafe1303c32829340437957f83356c837e2a"),
		expected:        decodeHexString(c, "f3269b27ab4f7eb1ff070eb2bb8794b4f22b0eec45c2ad1bbf2c5c0b7cd32d4de068f3e8282c3dcc35c9f469ef0a9d31756aedce00cfaeb309594b816491ccad"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_11(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "a22d4fac1266ce8f7cbaff6b9d47424a10f269e16b8c5278"),
		nonce:           decodeHexString(c, "7a04d4044533e983926023c5a8c0b10b"),
		personalization: decodeHexString(c, "a375cb70ba8006c74ccd06a9c8e41eadc445331e14b8e195f4a3fab0c1df6c96"),
		expected:        decodeHexString(c, "b3f9a39385616f6851beb7c37709ac1f30667df11c9d047dfcd5c7803e930bf384070e350f1d0b8faa39f29c98d9df59d35ffb520c8bc90be22fc77d4d67f134"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_12(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "48c4fcab285f5bb3edaa348c4506d236bb9fbb297dd4f328"),
		nonce:           decodeHexString(c, "8ada0b516dc49ac8c344f3b728274018"),
		personalization: decodeHexString(c, "528eed743085d3d416bd06984de244161bdee0bfb598a747f29c37fa9412976f"),
		expected:        decodeHexString(c, "2b473e1c19693b308def41f0cd99cdbf331c7d7d215cd5028e3619ad91c8e7d5e024804829f4b2905a33544996196d5cafabd2270d3cdd6689dd12462468ad50"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_13(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "607ff99facb269d5e85baf03dd93de5ea543168e6939d972"),
		nonce:           decodeHexString(c, "f849a0fa95d04b741e46d18f0612eeb7"),
		personalization: decodeHexString(c, "95fe53eb4db09aa85b9406530c9e994c695b86c9f04ee53f7f492c4d758202ad"),
		expected:        decodeHexString(c, "cae04ed642fa6dd8756a97b6a32494d6b95d5a5265412d27a0da9a70ad914ed31a674c6ed900be9afa493f8e141a08aba8dab75cba573c46637a9faf4c6a7161"),
	})
}

func (s *drbgSuite) TestCTR46_AES192_14(c *C) {
	s.testCTR46_AES192(c, &testData{
		entropyInput:    decodeHexString(c, "99e2df3bd55008e9cea0d0ea58c3022eff1d1e489cd6f8ee"),
		nonce:           decodeHexString(c, "b3907b8820c9f22979b663292f7ac788"),
		personalization: decodeHexString(c, "28beb66118cbce5eed825bdca9c8b65b6c6a498db6226e72e1fc01e71fd22c82"),
		expected:        decodeHexString(c, "43eb721f3669f918fd841d79e2a23648f88c9869260b624b8a3381bdc1d2b1296b6ba1e075cd4d3ee06a8b4f986c00bb2b9c5a1eb23b82ce081f51913e46a4dc"),
	})
}

func (s *drbgSuite) testCTR48_AES256(c *C, data *testData) {
	s.testCTR(c, 32, data)
}

func (s *drbgSuite) TestCTR48_AES256_0(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "36401940fa8b1fba91a1661f211d78a0b9389a74e5bccfece8d766af1a6d3b14"),
		nonce:           decodeHexString(c, "496f25b0f1301b4f501be30380a137eb"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5862eb38bd558dd978a696e6df164782ddd887e7e9a6c9f3f1fbafb78941b535a64912dfd224c6dc7454e5250b3d97165e16260c2faf1cc7735cb75fb4f07e1d"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_1(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "13199090a47fbd1984eb5fa9589345154699ef73f00cd62b07c34167c0327e53"),
		nonce:           decodeHexString(c, "5f968f93b659d8a5750a95345a8ae20c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d16878c5b06d7b6ced8e8aeb3a48d95ec8dd655733eec6ef473a8078dfdea600c0cc02168b4d6d744ee828ba5031941f8e3d96586407af79eba60d14af47d53a"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_2(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "d6ccf8c8143abfe5fd70626afc17f8aef172027c68c38f94ce59f7aed5e96657"),
		nonce:           decodeHexString(c, "2ebc66d2fd66b4bf1ed24faf744ffbc9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "6d474ba971a8339eca904a4c0dcf6265116fbc66cbe5dddfdc42104502eb210e3660e1b1b710b97d830c27212b33131d85d2f73f39760782f4b47d447ba6a68a"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_3(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "395d06b7549073c48252fb01f39542645600317220090029b2bac58a7a4c35df"),
		nonce:           decodeHexString(c, "5726b9911da8f166a84f82c06f53dc9e"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "0e1810b5044f28ef2cc7928bd632d7035bcdb9801e9d84f569a5b6d02d3cb5aac0a190bd58d6a08b6789529320c76817f27b7d331085346735ad371b5c9189cd"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_4(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "e502718e54c8a79f31529aba42404808e652477f595ab35bc54eaac7afaa228a"),
		nonce:           decodeHexString(c, "aee328ae82274d9dffdb2772315489b2"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "9d10baac91770e97be490db4d80d7007d6a20407813eee128acb161c6e36c225ebc42ca37b107f0430b69826add2e520c2f18fc07e32ec0a7b33463bcf48e576"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_5(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "63e60ff1739d4605f5c511fd0e3951dd3de657508a60d7c87ce95f39481a7553"),
		nonce:           decodeHexString(c, "b1c17bb34baf2c7f3b03b76e6897316f"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "08195061ded1ab84be7d5dcd630e7b903854f1284389e5e77849bba10c89083bb6f32f6815cfac45b7cad0dfb55498a47e875d21ed1842f58779845f2f24c6e2"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_6(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "63ec70792a6c78540e40db15c1b873fcde9bea1d741d88b81bd1fd66139df70c"),
		nonce:           decodeHexString(c, "fc5f902bdd0d7fe216b1423d4f41f12f"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "e366ae2988f6f37854859602319e64385244115004386a21eb25e69456c778947ae11c9532e5794e620460877a5bcd77f370e8a90d4be38206ca393edceae3ac"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_7(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "eadc5b6a6f8d10a0a7891703550ef40fc21ca500ee603a685ff3d1fb56ead70e"),
		nonce:           decodeHexString(c, "6c5a65d7e8fbc2a7cfd9fa7a5efbdcd7"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d428a298cbfdf20ef7bee8913a26c53eea49933a2dde421bca4b1c6b86506c6ca0ef0bfa13c030505748d5737b232480edc0bbc9e7b7f3ff8aefc292106a6254"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_8(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "0f9106588bc927eca285e05c8c170a68e9a103102a106781d8fb0a7688e491c2"),
		nonce:           decodeHexString(c, "2e58a859b7bdc816fb40b9ee5916a925"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "e3decd0810e1db9b77490cbe78cad252c261f0f38051e5ba1e1ff9852d0a68a1fbdc798eb196f2ba132e6a451dfe6a9888833666bbc304c617c2d610bd5e489e"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_9(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "4162a42cb0d83d3384a1dcb9dadcc6d070ebd1683b3c0900f31d7cbe050e6aca"),
		nonce:           decodeHexString(c, "15f154ac7c825258de3d6689fb9ab46f"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "3eadb8acf7de3bd098297067366183a07990550adead0bbbe48fe87bfde21ca878e3cec77379bb884f28546c57c40e28e723f8cb8c9e04a850a46dc75fe90368"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_10(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "3652201d2ced056c83667157d7b0ed02dc4ba3f6d43799fda4a6d5204c4e3bb4"),
		nonce:           decodeHexString(c, "37f181837052d729a70b462f5c1fecc2"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "b5bc59138df351490057cf9df97e991745e03b30ee9684e61812d8453dac5bcf54996d95ca1d71fbc5992d18de9bf224f2ba42d24e3f8d13e341182274cbcd32"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_11(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "c1d3cba48d328ad2e41c75d21cb53b69ff16cafb51c241d36e1d752fa3de2dd3"),
		nonce:           decodeHexString(c, "b1b6e9d1a50d4fa654ebd44663a8ec01"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "2b9ba1a9bc56c7f3e5b2d587aaf98d1e7da566cff5f2e7f91710d43f8cf7e485632ff451110af873e0b33f0dff468f09c5e3233d3807cd739982b5e5885434c9"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_12(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "d6702ed5d03596d2d1daf9ffe9c0a19a479477f7e914654f7b114d18b63f507c"),
		nonce:           decodeHexString(c, "53b29d4b42d3c8798dc1e63df9f30bbf"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "fb4e3f7b0c6efae148130f82d59cbc993812412e0f57b206ff0864f4dbb8e407eb7de2c68af8a5141610ed4765274935e6f37ef4eed37de273788576e7b89daa"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_13(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "7ee61faa3b3fd983ba3b350bee25bd075e2e6f172be964bbe4ff210fc92de651"),
		nonce:           decodeHexString(c, "a0c4ebcc0aef848ed7b00d3b612972c9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "a2a190f0fdfaf66314d4e0071d4ec3d2914fa93f8964ef5608ec386bba034670ec0177350a780214bac579290918fba9b21ef9f4f1210034d0809fedc44d33b9"),
	})
}

func (s *drbgSuite) TestCTR48_AES256_14(c *C) {
	s.testCTR48_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "86e3ccf09382f46192d8937eba1dfc04ad8e12ef196d6d941970b16e15438bb2"),
		nonce:           decodeHexString(c, "e2f9cc7eab26491e3eab3cfdfaadda90"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d8bc7e8220cbe079b4403d6b47e7976605735f953aabf93fd0039396ddb4088e436052722c4325f90eb868512ae7ce3927370940b200f3cc81da774779e90d7c"),
	})
}

func (s *drbgSuite) testCTR50_AES256(c *C, data *testData) {
	s.testCTR(c, 32, data)
}

func (s *drbgSuite) TestCTR50_AES256_0(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "5416e77b5e1d872d4ff91973b1be66bc07f4a99e30db7d0006da006fcfb082db"),
		nonce:           decodeHexString(c, "7a811ce62b9fd34af186b2b3e50eaf5d"),
		personalization: decodeHexString(c, "71ee0c7699ac0e805632f2058de38bf872b8340f89998f7a8a2ad4ac045ae6ef"),
		expected:        decodeHexString(c, "68f5859cf76f94c445d9fcd34fc17ac224c3d7d7c2fc38faaf3c24be6cd3cd93b7f9d8a6146f5ac83ac1d7b1b2b7e7ecbc1a2e38760ef86a577d402d85990d9b"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_1(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "708eca2e3a9265a790607edbe05fe342663f84c6617eda14f25276a943901fda"),
		nonce:           decodeHexString(c, "75afb49a184b23506be14926cd4a03f0"),
		personalization: decodeHexString(c, "cbb48ef84146c10e02240d8740d3487b6a4208405383c01a664ec7d3ada07e2d"),
		expected:        decodeHexString(c, "26b0aa6e822c4cc912cf1dbae669c7dad0bdcff65f22813afd06225b7ff799f7803b3ad48bc88d2be0f5a357f620cc617f446fc6d212592ada69b7dc8ff4a222"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_2(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "44cc6b4433cec615c3c214e166c7dcff258f8cfe5748e642321cda2f7db426e3"),
		nonce:           decodeHexString(c, "6a2526954b5df989d61e1faf93dda2ae"),
		personalization: decodeHexString(c, "88226313c7f1ec03cde377970c8ea7d741a9f21a8f54b6b97043bc3e8da40b1e"),
		expected:        decodeHexString(c, "c1956c4195adfc3ee71582ab2c63edc0a78af49ecc23a3ddbcf2dfaf80c761fd6343af6d14310e719d8cd3c6bbb491c8690a7dd8a168cd8a480217e5dd2088af"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_3(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "54ccb1e5f044447dce52a470f47fe2682717dd296d64491ee2acc99e9ad6566f"),
		nonce:           decodeHexString(c, "ff4cd3185611cbe06784e32580b2f23c"),
		personalization: decodeHexString(c, "132ef542f907b84c443d1973b3909b6d9a0d9124d38bd1e7c8833f48aecfc08d"),
		expected:        decodeHexString(c, "0f2f56ea8b911cbe59a7b8fab1c710a7ebb6ec9a09555ca49ccddd9afd38ed61b855cf3f33f2c5bc616df6cb1726968483c69c1849e0f1b46ba029aa6f5debdb"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_4(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "3d3fdd9d90acbcee07002f17370045feb5eaa334fd74594e112114d3928dd5d9"),
		nonce:           decodeHexString(c, "f85095294ebc5fddf44941be5ffaf10c"),
		personalization: decodeHexString(c, "59e2ffa164733ff11b5a95eb99a78366906de4fac64e512400081116acce5390"),
		expected:        decodeHexString(c, "1244e532799f1ea4ed321894dac51b3c78d2fa5f0e1c922ffd2ff608275400834d03454942d31a2014ccfe07c2354112363c60f48dd12b29a3734128a59bca21"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_5(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "2536d7ea906d1eeea83e1c9da08cf76d095b7d4f55b433bbd5225aa870cc3f50"),
		nonce:           decodeHexString(c, "17704a7c50f937acf626e854e0b95f80"),
		personalization: decodeHexString(c, "b61fccd40d9cebc4b927ba773b932957c299f779c4266a83f169b6de507d3127"),
		expected:        decodeHexString(c, "8c3b0b71d6432e1022976c9c8ebb1406b1da995e2937221d18d751816825fca064534e2169cc63b5070529ff02b5ee5b7081a08ebdd87862595ea37a95c1e4a7"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_6(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "fee982d87df4c32057cb499df6ea21dd6916dd38850a872abe34360f2914b56c"),
		nonce:           decodeHexString(c, "5a1ff0692d9adc5edb1af77461fd0531"),
		personalization: decodeHexString(c, "1c0d185fa4bededb781b8b73616def7440822172d01a619a2066b79eface685f"),
		expected:        decodeHexString(c, "d7c228e7002bc2a1a23a16e489ba8f0a61627ab4b2444f00f23ed3717b87944feebb169ee4e979089c35e83f243079e4771918862e15036ec6a9c9eee4bb00b5"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_7(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "25327b05f86e5079ab552e987ec7e2816259510933bc0c7a05c35e163b47b4ce"),
		nonce:           decodeHexString(c, "7ef8d6308ee6b5061cbcb850485a15a7"),
		personalization: decodeHexString(c, "9a39ed86b2e8290f74ca886b320d8e23d8a1d4953c2762f81f071d842ce951e3"),
		expected:        decodeHexString(c, "211971632f8e6c6082f11076bb707ec9c65e3434c2727e59a5d13e543562a1da4d7bc2e483035d4668536ad1d69f6119045999524d92c48c4aede622ee53638d"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_8(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "7b163dc11e156b1bb170c22dcdd121eaea301d6172e3c0b4bbc4f377be040940"),
		nonce:           decodeHexString(c, "3233f156372869d937f3694c50ea9832"),
		personalization: decodeHexString(c, "7d61daf237bdea1ba3b0301f02cbe61ceaa188eef85b95842432470c65a13367"),
		expected:        decodeHexString(c, "8b9e1b2aa950d016aee55c3abf0c7f1d3c1c0d8a62ce5d59e84b72a1322f1746efaf0e46e67b176643c7b57e3dbffbdf32ce7afffd6ff7f23081ae4c4f973eb0"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_9(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "96cb5f531e02ac254445673de5e1dcad51761026e778a03321b49c597745731b"),
		nonce:           decodeHexString(c, "139198de0ca982c49b8a88e5ab886cb8"),
		personalization: decodeHexString(c, "006f172e5186fb479f3cd8e3425a752b8f8d56a3f176c6faba6ca4a4ddeae637"),
		expected:        decodeHexString(c, "388e5c98c054393084a94415e1a9ba32044c874562c185399b969994b6406f9ef86319102fd89df670903567f86de8676b0f72298aa00191151262e7e0898791"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_10(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "585af27f7f524ab4c5a7fca43b70b464cd71d7c5bfb4455e6dbd102d89335e89"),
		nonce:           decodeHexString(c, "0ac9cb4d14bf522616c613068698b16b"),
		personalization: decodeHexString(c, "0ca0ed3027d3b47f783feef18a40340cc5851b342f39f2b7d908910b96792e3a"),
		expected:        decodeHexString(c, "f4c2b3b86471b2fb446992791942156d85fac3c937a8e50e79aca792c6695092d67f807eb7ac0daee31d68f66a8348c33c315e4be1ef586dbd408e63db1d8ab0"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_11(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "4e7b5b7b30b8b90b1bdc043c788054584fd3945f45ca54709f89c19b75566dc3"),
		nonce:           decodeHexString(c, "85a400116eb9030c9249ceaee8dec9c7"),
		personalization: decodeHexString(c, "091f590f584230696163b6b53a11f52c2e9259aa0982d9897e7dc594199d48cb"),
		expected:        decodeHexString(c, "c89bc08d62335d3c28464ee2762285ceb0eefc9577195a8f89d63e05520201823af887305cff9b2d7fd17f60b5117fdcd0ed3ef18bac7e5f15080802a8ab139e"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_12(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "ac2dd682a96ee0e818829db562fbc30a7c00e0193147c533e2a51b374a467958"),
		nonce:           decodeHexString(c, "b8e6396449d884e9c6ee7ced8c7ad92e"),
		personalization: decodeHexString(c, "1087ed01b4a853efd8b9d3adc5893936499dd29745cdc6d3d893adc81d5a1bcd"),
		expected:        decodeHexString(c, "e1154b376b8c7c26151b7f68a5db274d73101803c1e7c2f6d10f374e4b883b10305a5d587676020d52c6c4a7e553bb96a953b6e9f2e4a7dd60ed4f94d123e4a3"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_13(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "4044299a1a97126a524597c8df580c9f54adb4143a8282336496d9ed41881357"),
		nonce:           decodeHexString(c, "787da4f01598b6f09e3a7b6865f90987"),
		personalization: decodeHexString(c, "c778254d48628e48f4dcac3c96dfb2cf8410f9cdb6c00d625b411a0147bfe16d"),
		expected:        decodeHexString(c, "1bfab93a54b7a0bbc6a77d63d8132a5d6c191c84e09a8591d2b7d88f339cad29d43cc8d0bf01d6ea9ffe27fa1ff7699c871a8411332ab0bcebce017aded3efb6"),
	})
}

func (s *drbgSuite) TestCTR50_AES256_14(c *C) {
	s.testCTR50_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "56e5dd70782fca50493b5d222b83b6dd7bdf1022d304f4f292db385f3acd8ab0"),
		nonce:           decodeHexString(c, "f0ec46de4a9aab3f9e7fd47f60d20450"),
		personalization: decodeHexString(c, "cc542e50fc0abb15b7dde7b0dfe841c79f748244a0d1d0a02e9f908ab99c61c2"),
		expected:        decodeHexString(c, "7427c7018c51f3b060cbd8fe253a8b62b672a31fb848f55831be1b045f6085c91364853c7e353be0a52dcadb30928e94e1772bedbf27101a2a8ff8713d9c9f4c"),
	})
}

func (s *drbgSuite) testCTR52_AES256(c *C, data *testData) {
	s.testCTR(c, 32, data)
}

func (s *drbgSuite) TestCTR52_AES256_0(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "8b0bcb3f932170416739ea42e7dcdc6fa960645bc018820134f714b3c6912b56"),
		nonce:           decodeHexString(c, "bac0fdc0c417aa269bbdea77e928f9f8"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d9c4fd81f6621a8cf06d612e9a84b80fa13d098dceaf2c083dc81cd80caedd105c7f2789963a167d72f76e81178001fd93de4623c260fe9eebced89f7b4b047a"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_1(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "67b6e84d5a560af4d92745853da83c4e8dcff469869eca69981055ba4c6f84c3"),
		nonce:           decodeHexString(c, "aabc8d3ab593dbea35fab1ff6cdc26fb"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "e74ad622a71298983aa21066d788fdcd6afdc9aaf7fc8a55534ec0917d6840d15c1ba2f0a703f04b148bd7bc4983b279a414e3937c17a8181e644ea0662dbebc"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_2(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "be57cf16b26481aab3164b8060c29f17982711b451188deacdc9805ef7e016be"),
		nonce:           decodeHexString(c, "85484daa20b8602507b3d76850939e59"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "4b95469216037de3afb2790dd1523473cb8dbdf7230b0f8543f738c6baa0a1a46d13366fe3164f245676dfe1af0214c5581e82790fde30b0203e4554804b9826"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_3(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "3cbbdc1bbb6f005897d65384ed7979df6d7108559c3e7619d6dfc8dad8e6549f"),
		nonce:           decodeHexString(c, "7c9f78b7d15ea73244123ffdb4489f0d"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "c21af26518c97d6763b753939ab0df09cd2d59fa1090933e2641c9439f79fb3b9022fa6e07c9950ce9eadc3327dc49f91dc8006c7da52b021e5ae83582f30475"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_4(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "06311cc81e75b12269bc67ed0e1313480f324b752a1fd783ce09770d1d4000fb"),
		nonce:           decodeHexString(c, "ae2cc517b215855b1e91d1f7942d931e"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5ed914438a62cc4d463367b17c89604290e9723522ef801ead515ca352e099d6fa1362ded32a3fa36b2453422f114f8cf246c18c8cdd786aa243801caef2b2ed"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_5(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "5927ad80b12b86269eb31e85c3b122a7b82d8ed5782276344be3f52edf6280db"),
		nonce:           decodeHexString(c, "e6740d3713da617baa5a5cd5b4d6bed8"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "52f310ec9adcf491de58dc52595ddbb1fee966e69e60ba0e65bed7fbf8dc3d58148538a1b456aece1b41479dea06e1015b7b10deb66c9774029f0b52ceb1f58a"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_6(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "6ec135792f27026772739e893ac5085220351f590575feb9a3d3792ac913d8a4"),
		nonce:           decodeHexString(c, "a0969cdb40591436cb637fe093156f16"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "08d2735718f9c882322fdef7b121339d5a7f59b074d3c6f8b8fc517a072d356bb702d90dbcec40e01c4ddb6281387e731b2bd0e72a6164ef4a721dad01023b2a"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_7(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "e9c48be6d65c8dd3bb372231dfa8a7078df501a9238d631161cee3f0afa97889"),
		nonce:           decodeHexString(c, "1d32b5a37be3ad6096e80c7fabf8a79b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "024fa95b3cef85dd9a227c93a8223b1d02d57fd506ed42a103f7ee0d9205cad9fc51bcbdb7c401d15fcb06e4a5eda17536fe9e93812eb99a879b9106ca0b6d8d"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_8(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "6f37bb62e2b301de8ff6c5a80ad8634ac718731ba789b247c3507216f64bae42"),
		nonce:           decodeHexString(c, "41e514980d2c79572584d0e60f202c67"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "c1ad074fb5748f09a34b8f8b8b15789d26a6aff4714095c54f62611d88ee2a45a4ac1110a3e4ec9d59ba85ec1d80daacddae3ea58f9539219afaf57fe953536e"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_9(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "54902f1ce788fddb8d815a80e5b02da36ae09fe2cecb6cad766b757238a09546"),
		nonce:           decodeHexString(c, "52fc08265f84bb3b2f5dab01eb8ade3c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "0d114d27d1c4ae86186c313e4f1bfb002de3f2a827b7153941f90f994b56c8c266816cc69e1072a2b3be19c29f80ab3d039ed5722fedcc016fa82b3961782858"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_10(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "fe55e2f96b16d2dd915ca11e2504d720f1aa1918631c472cb0b4499b7d886b56"),
		nonce:           decodeHexString(c, "8530fc3eba5a01b763089e8043dfdf79"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "be3e72b3aa0b07d0756645a5f5a74aa6a015b7cea68581fc8ed762660d9e54a7810e712d81988ea706ad2b938e33651a3e8a6755d592f5217ff2c9ea8f41a5c4"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_11(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "f5fc14971b966a4f95bb13d55a682bfdc7da7f26c943db401a179f23f47ec3be"),
		nonce:           decodeHexString(c, "f051eebabce3599021c364811b11f28b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "9e184c66834d73b4610e5362cdf908a5f84939ecf01a58ebe6c2dc50ca0c24bed994a82eec998e18afb95727c42786f5e94d60c606576af83f3b335279a1014d"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_12(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "26e6f44b112f5240f0e2249c2aa4d5df3df09431296feead01cf659c62de5007"),
		nonce:           decodeHexString(c, "8caa8a24c28a17cfb68613e2fb3bf37e"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "91a3ad60da8a72d70405a1ba819ecb147f338868a41668db59fffb0311a8344ae0ea58d97b6d3c316490b5c623653e816cc1b91d5fcd9b4d70e7717d6ee2b702"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_13(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "f291d7e5c2961313184f8023277b8fac09a543a268ee506a53417067a8cc12b6"),
		nonce:           decodeHexString(c, "55243cc5460e24bf0c2f92c0cfb23eb6"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "809d31a2a26c940b3f045f16b859f2708fd9196a3983514104d80f1f68166af0e48b98271b76a267ea1e05250a3f7f2581c1d6356708c293f208944b3da0a425"),
	})
}

func (s *drbgSuite) TestCTR52_AES256_14(c *C) {
	s.testCTR52_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "8b0878d733c8a48a79fa674dfa889b34fd0522016d2f68fddd3dd05f1af78506"),
		nonce:           decodeHexString(c, "f772496334fd93b46647297ac444629c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "9191c5970a37fdd5715a7c6a09a7dddbfbee8d37ee45a035b33823c6e86fe0e2e2e57ff7d08e0a4f39d81c18b9b5fd487e84ca95df372d36dcbfac682e044b34"),
	})
}

func (s *drbgSuite) testCTR54_AES256(c *C, data *testData) {
	s.testCTR(c, 32, data)
}

func (s *drbgSuite) TestCTR54_AES256_0(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "8782d516ab2e0720816d31e841c4976583f5f2356d4a6b75baa0c854d81e87df"),
		nonce:           decodeHexString(c, "d3a0df6e410cba3af82b2e914e52b19a"),
		personalization: decodeHexString(c, "9460e6673c94ac44f812673c25b8905456c32fa7a88d019c9b9af0e9e6dfde32"),
		expected:        decodeHexString(c, "73be5aca786c4d2001f026a48fc32e0d5b9c43f5581589809f103cf91fdc33aa000703c5b9a7391c4c75126ba00f9f9cf368b0f92a72905ec11f670244d02e33"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_1(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "a3a058ec8f4963e3e4a5e7aeadead48e48a130f04ae6785c184d76ff8c78134e"),
		nonce:           decodeHexString(c, "ca4ff0c8c05db6d766f356216c3b5fb4"),
		personalization: decodeHexString(c, "cf95338ce69272324c751759566e99eb9a2a618cedeea977c360a35be7db807c"),
		expected:        decodeHexString(c, "f593fecdecfd70d9f7cc093b4cf0502f178c9997ce7f3b95cbafbaf6e575637d344e2c9b7ebcb9ed6048650639ea48d321c626086b28002d863cafede091e7e5"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_2(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "91f74d6c798f6e1842e36aa61019682e246a2eec04aac8f7c5e849dbd6fa677d"),
		nonce:           decodeHexString(c, "800723008b744351979ae85d92fd217f"),
		personalization: decodeHexString(c, "c9b38f9b98b7a0043b13d1926c27265521f01316e8fe79d2efa8b817b23aafc5"),
		expected:        decodeHexString(c, "5a63770230a4a048645ce1a75e50ea792c219634565f24ec52fae6046506c5b0529a798c6bb71619a24bbd71f90335e93c41de3fd0fd1f3ee3204b9c6064b735"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_3(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "aa40fd98eea752b731545a6b9386b2ff356ef7d9ce88daa2219a5c5fe57109c8"),
		nonce:           decodeHexString(c, "110acdc86c06edcf8d612a4f2df6ea72"),
		personalization: decodeHexString(c, "f17a8e0d460e758747c461782aee6dad4a2ea8cc26c33b34e797c9ae8f8fc632"),
		expected:        decodeHexString(c, "3da63f9fee355441fac4084bb9766851cb6c60b7a94842d2c7623b807a4620ccae174810c21d92d5f2676f9e84a5c98b9a8a23adf72ddfe39fb788f289217187"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_4(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "bd5daa18dff57e90762dffd35a05a4a739ce7ebf087f4293f2c7a031d17df9eb"),
		nonce:           decodeHexString(c, "af62923ef214462418439ec8dc553c9b"),
		personalization: decodeHexString(c, "fc9e138be9170c8312288191039b033bf41ef1d47f4e642357866b875c7f183e"),
		expected:        decodeHexString(c, "d40ecb4e47e55460c40047d60f852878b915268a4f13796cf5d9aa0d67f6da8809847468d7e04c039a9f9e3d9e5b4d53ce8f66fe7d88a4983c5111cef6037b33"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_5(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "5a57109bf7dee009922623cb8dd7c6029d14e6f3997e9a5405739bd85027d31e"),
		nonce:           decodeHexString(c, "a4c9511f90665b44218d45c627d8be7d"),
		personalization: decodeHexString(c, "2a6be524c3f7a580150f699bcc6ba10b1e26cb85712621ce00da3f363c8b1c46"),
		expected:        decodeHexString(c, "aaf4c6e3ada51a1ae62c24381b21ad4aed83d8ae3945eb71938c1ce2d58627115e5efc3c58e8056e7f7190ae550a8f1b46d1b58ae3ea9af0b7b1081aa7fd3630"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_6(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "f43d1631397bcfd6c6b385106fef8e72f97284502862d6464e610cede7b9d8ae"),
		nonce:           decodeHexString(c, "741ccd9aa487fd775378eaf31eb877d8"),
		personalization: decodeHexString(c, "9607f84f0b2237626da9250bcba646a06dd80049f9a28c4f0b2b631ffe730d5c"),
		expected:        decodeHexString(c, "5097a9f27eb3ae423281b93b77208258150342ea4245db7f1c4e7347b6a380c81e27b0eb00610e842e2ab57ddc0345676719b52af6630d5290975e6575a017b3"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_7(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "b9f8f6f67350b2b674a98ec3ccf9f94d5eb9ed74c674ec305dd98fc5d8c64d93"),
		nonce:           decodeHexString(c, "ed55a0e682bac6ad5d6f86eb31f3b15e"),
		personalization: decodeHexString(c, "9c299b3dc0782f0746ee9aa19c24fceeb7b350dd3de9727a19708f41d99c86aa"),
		expected:        decodeHexString(c, "2ce4af0ce1d65a8378d3b896ad2e16b9e8bd2a92d595c107159c386a59937054636b6425ec731c1d838c530153086468e752ea34d2ad64c068f0015341aea6a1"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_8(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "c699205859904c44a7829c81a581d636a97c3f5721125769d8070b638fab07b1"),
		nonce:           decodeHexString(c, "4ddbba6a5a137e24ce4b283a8ff527e3"),
		personalization: decodeHexString(c, "22dc486c4fe19af1f82089bab3d01ef160b6f5e5948369c0acacd6a5f411f737"),
		expected:        decodeHexString(c, "ed8eab6c640e6e061725d718f767ca510849542137da84700bdcacadf1cae5dd03f6edbac0053f1b7fe2cfd5363a2fb8129e149ac74764470b8623474e8b204d"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_9(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "164af0b91ba367e7d949e0687a7cf1825919745a2ff41a28c06e9d647cea545a"),
		nonce:           decodeHexString(c, "a610a3cfc51cd6ad869fa85086124aa7"),
		personalization: decodeHexString(c, "5d03586dcda1f022dea643a61fbc9e25c7f07fc15ba554d5adb8d4785ca76bc7"),
		expected:        decodeHexString(c, "a71aa159b481353e1ea8a2fb9f0fa6fbcaac1c24104717f9894535f087e29079b1c3168fb745e0503101eda84842980f9191727ea47ffb570cfbaaedff055d9c"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_10(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "6a591557df6afe71cad5acffbbf758f6829ea887559bc1c3ab6f1ddbef928b0f"),
		nonce:           decodeHexString(c, "c872c4f0d9afef9be408e0ac48878bae"),
		personalization: decodeHexString(c, "dfa1c2db43f6129b31fb4b3266b17d972bb467957261df3a67678794497601b6"),
		expected:        decodeHexString(c, "44f4ad62710d87a109910908a4dd1a579151654b7c440686f903b213285ea886b72605d8e6af0c2cb9286f46b87d6cdfe1702481248a816e887d766858b221f4"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_11(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "e2bc6e865abeb1bbf13ebed7e123ffa32687625c67ef561de9c2fc6f14efdd77"),
		nonce:           decodeHexString(c, "8e675bf0ac4b6b1fd43a5ba3d55eb5fe"),
		personalization: decodeHexString(c, "2ad79b01b477c2646be5761272330d98405edd775891f619c9666791dd2a9b68"),
		expected:        decodeHexString(c, "f481f0f81f46feef263410ac39ef51f1faaacf22541de8e4f91452b690bb5c6a211dc8751dcf8bdbf4dcd3554894bb6644fb7995b40c04ea4c3a979e97b5bcfb"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_12(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "ffed524b1d1f1c5cee98762858387d1a403245ebfafec0c0a40ad815db9c83f2"),
		nonce:           decodeHexString(c, "95f9d8655211a0f80ef8e376abd8f171"),
		personalization: decodeHexString(c, "38b2ef1fcf66ab0ad6f14ef68b7db36a7311ff9964ff4bd501c447fd04167340"),
		expected:        decodeHexString(c, "2ed810830b5c7bb287eea85cd32ede0fe070049a898e0bbd7a6949efe3bd3c19075acf87ff277144abc2e23ceb416b705b7a2be2fb0a0dc9164786f3527cfe14"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_13(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "2a92903d4d3aba0916cbc072f98a1705446a126d3e684d0d271b78a46ae85c6e"),
		nonce:           decodeHexString(c, "2c94361bdb040875a64e1f21b9ca8550"),
		personalization: decodeHexString(c, "353f18f9c331a1f77a1b0aa5f970762ebacbacb4273a941114a3c783b6e4088e"),
		expected:        decodeHexString(c, "f9e3f3317725f96b137baa2ee216d70cc4b45c380520f81d6fde1368ef701087a29098933a38ee757b9f70b182e7063e3630a43bdc1f21c30c4a47bf7fbb98e5"),
	})
}

func (s *drbgSuite) TestCTR54_AES256_14(c *C) {
	s.testCTR54_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "c27529d8650a498fe1aacdf12d8a56c73ed1c4b566fcfc57b43eeb2f8c1790ad"),
		nonce:           decodeHexString(c, "0945153fe4789cd0a74510dfcda5de71"),
		personalization: decodeHexString(c, "cd98bf629da5dec86d823bbd720e273fa09b75a304a2c3139c38386bb4468553"),
		expected:        decodeHexString(c, "ba34d46c29afc91ec34cad8485b5dd6e2fb580a217859f7049a6b23eb93e66e429ddff22cc647159ef1195bb67c40cb33883ebd8a6b9f71ab0acb89de774a593"),
	})
}

func (s *drbgSuite) testCTR56_AES256(c *C, data *testData) {
	s.testCTR(c, 32, data)
}

func (s *drbgSuite) TestCTR56_AES256_0(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "58a5f79da44b9f23a98a39352972ad16031fe13637bd18d6cb6c9f5269d8e240"),
		nonce:           decodeHexString(c, "aaa46610681167ff8d4d2c51e77911d4"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "c1714f89459ce746b151509e5066d4811a06ad06c1e9b13b50c0fc7cdd77ceedc233908ebe1ea8140ec2dc262a43201be667008e081e5476b19b27214111d325"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_1(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "a943e809630413de6207746d0d0341913f466af0ae893cfb3406570b2fb791cf"),
		nonce:           decodeHexString(c, "907b9cf7f9edf04fcf3510315dd0c381"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "c3b84888c5244cdcbb946299cc3848c379a9b780e21f029f0bb2fe815a2d039dd7aa8a2e808c2ac47b8a9cb6860b970440049a65d815e3369ed833c76124aac1"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_2(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "4d43ceb1ce9cacf56403a0c9905daa67a2acddd0e4be6a334b8c4434f4c60455"),
		nonce:           decodeHexString(c, "9772aaea3cd30ca776d674bcfb884e18"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "85c87c4715c16c4b7528d04023056398ff2828e0b649fbd10a297b74fc3dc0df4966bcfcd4f82fdb228faf102d52cca0d3ae8af7f0c5b30fff62d0c545d3de79"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_3(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "d410cf13cae365faf3172fb0c2368401f443e789a62b3ce6bc40023249fe7dee"),
		nonce:           decodeHexString(c, "22c312b52a0692eb38763332b6cd4ae9"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "a005f40fdefa1e1dce22f735bf3e87eafc8ec3584a6b2b8045be53bde0b1cdc46be380be860538ca0e976eeddae4add2204262350d5f6e19e34db0fc47dcd0ec"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_4(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "f80972a5cd4e2e14b1f5214dd93c549dc51edb97c1447d52f3e91b30c15b748c"),
		nonce:           decodeHexString(c, "e1dfe4aa777a0bebd1fe936635a5193d"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "f389defd88ce73a6a5d81a32fedf26e005a5d42f7868fba40ddf20df6325fe34738da3cebb62b602217247fef77837fc73dbef33b813b26eb06be2ad05069882"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_5(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "92ae363fcbe35fd5f606d21d2094f95b90e6b11c21494fe0ea3f7e8d0ccb2bc3"),
		nonce:           decodeHexString(c, "88b31b931f49f86c430571af1a352a3c"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "4f4a1e0ea821bb8add74afe6e12659813338ede3d1d3535808342a0cb3d066ec13a8d52997336607297e56ff53ca5a50147590232d26f041f76ff9b5823f36f4"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_6(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "01ab1c4096c4ed6d241074c6e8b3e4abe7935235d6f48e913b780d219a71fe2d"),
		nonce:           decodeHexString(c, "472e39d4800ef0eece7bed9544d5af5a"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "bf7acb20907e0ba141e46519aa31d46ad45da46e014b2420d48d309cb0d4703e82c6a78e835efabb1723b75b8dace5968042416075cb71b9c149806cc4214758"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_7(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "f63021bc4922305c8267b2cef767e3bbdab72e119bc8317ad8fe2c27694f9e00"),
		nonce:           decodeHexString(c, "e6ccef20f87519d5d67716e12e1bd760"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d7eec51c186f5990bc3f55db342a41a2ea155034af651c3c291310241d4655733cc9d781d6886138b3f9ac595b956f94daff64dfde94d3c8f2b81182b87eb832"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_8(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "a8fdae7e1116122d2b42d154a6329aee99ae6f3bd65180a03828a3544b80f99b"),
		nonce:           decodeHexString(c, "10b14f568e42ff31cf73b1f3943aecb8"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "4abde8199c0776ace8afba196d2befc6d6888f3e9560108895689017ac654a555722c0f14c53b19f3cf671a88c10a915ea720e7dc9b9e5e031915e330ebb0e4c"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_9(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "9aefe43ac5490defc5baadb12e66838de9ddafc944cd686ecb059635403452af"),
		nonce:           decodeHexString(c, "e9256f4902c7d7f6bbb1dcc156ddf9bd"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "d96d90ac5acab54be18da19c7615b4ce471a22416ffe2043e782beb885269bbbace4c416050cb4fb90158fda6da58bf8660ef9160d1c15fdd4c067c5c98ccdb7"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_10(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "3618b56386b7ae61dfc184999cc4587cf07eda8be861f1c64a28ecb2875890d3"),
		nonce:           decodeHexString(c, "deb03acd0ea6f0591580b8afbc2ff880"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "f605305d87f209ac293ab9c8c1d3128418fcc9175aec2565265b1dac75e7aa8057953e0f6854800f37ce405eeb24927ab78e94fb694b668caaeba5271fc486f4"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_11(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "c09cbd478e38c7520cd92432b5fce8c15758ef56c4e457fd4f2c97d52e8f4566"),
		nonce:           decodeHexString(c, "053ea292367476415df7fedbf25d6954"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "00a81e0b7db15698b89a75f56175f449dbc4211ba354b46f2474acb55dc170bf6c0ee181dceac3daf5fc81a57c63ddc281cd24cd6e7f25a1b40adb4d27c71319"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_12(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "59ef31ed6854016ab4ab7c4af72f5b16ad2d7b57a18f79f8f29de26ee4200961"),
		nonce:           decodeHexString(c, "8e8d6f8378634314f72ae057991a333e"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "1acf343fca1f42d47a1131e5d18906b5b73fff61461237882a58bc5add7527a5dd03d2b1967b70226d1119ae1435911515e3eaa67acf144214712ff3346490f2"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_13(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "04138e8fe8fef7feeab54ee3e5802e027b1f394487cac36d1bb0f052a7a1075c"),
		nonce:           decodeHexString(c, "d6b6484dbb46389edfc1ac969e462214"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "09f2c686d9b01d49f12348f12ff832db4ea9105ea6a148fcaaea32490d2d843a282cdc30c011191f8ba77c224c297012e261c8a739d0777298eb591c99028ca5"),
	})
}

func (s *drbgSuite) TestCTR56_AES256_14(c *C) {
	s.testCTR56_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "11b1a307c266d407c3ab12be774ee31cdc918ae14cd98fcf02c3acc882ddb82c"),
		nonce:           decodeHexString(c, "e4efdaf6f01eb5475d82c94e129a35e1"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "6a5263644a101e9f03a40260a68fa7271a399ba954fee230613a18bc18cd5c5403d679c7735cb1671b174e60ed828e707377f55178bb538bf4267453e5c85b97"),
	})
}

func (s *drbgSuite) testCTR58_AES256(c *C, data *testData) {
	s.testCTR(c, 32, data)
}

func (s *drbgSuite) TestCTR58_AES256_0(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "47d570e3a0a20c0a2010673a6131dcc3202679eb06f3c1b82a710edf92d3774b"),
		nonce:           decodeHexString(c, "881fa7fdfb70d106b9b9eb0254b0eb6b"),
		personalization: decodeHexString(c, "0f66787ef9b90364517e3151b158becd9df4060cd92ec88da3a6dd7b3b18e54f"),
		expected:        decodeHexString(c, "e64e8dcac95e0e46f5e6c5571d077b574b1eabe4880bbc0bab8e08e2148051441165c305fc09d60765190346af27a0df815653e81f782ab7fee55dad23ec51d1"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_1(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "21cf7b1f014995ffe7fe84543f3e9a75cb3f99851cf21c4abbdc387330d5c7e9"),
		nonce:           decodeHexString(c, "49a6eea4263ee1f5d461907dc58b44fb"),
		personalization: decodeHexString(c, "14d53975f852bcc9a1c5ec9f4825a04721ecfd87f2adef099a5b88e27d777b03"),
		expected:        decodeHexString(c, "a26c9905c9ae138d948be73c4271e7e0daa23161bc6595154881ae6053599a21aa97e57f3ce34d30f69647e970e7827039932615d970b47575964ceb8f7a437d"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_2(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "179e22e2f495ba9ae352b93c836b6933e28a2a184f8982c04e25e7eee66f9f7c"),
		nonce:           decodeHexString(c, "c184e842d2555e56888b7b75189e7775"),
		personalization: decodeHexString(c, "0fc74e50a0fda79bb31d5ebb308aa97ccd6e6f17dcec14976f4e6b15ca1be341"),
		expected:        decodeHexString(c, "47497647111744c73dd2d05cf57d65ceae22e22cac44b098189c5a5c2f781b74669e6d669e38ea8e5b4660e04c0401c4a4e64c331d796d19b7350a6a3e4619fc"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_3(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "60bf6d9573ea9398074c3d6e04e0e822f0ee95b67dd255598812e5336acc2336"),
		nonce:           decodeHexString(c, "ad27d7b274f3a2189d27bc547d6ac410"),
		personalization: decodeHexString(c, "58035bae9ba67b890b892e3a974e331d99d15c607593ed21a51dec71034d142c"),
		expected:        decodeHexString(c, "518dfea250e179e2738283de6bc29a069cd9f6793ce11ec801ab2d323914e13455876c4db8328a5b9a78fd061cad66232d1057e5d6a385a297adca9cdca57caa"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_4(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "d39fb356f66ceab81c4cdad89ee2e6657f8b36e845d6b8b7530a6188c4c26a8c"),
		nonce:           decodeHexString(c, "b10922b4a3086e4a484b5039996a54b4"),
		personalization: decodeHexString(c, "04df5d29dc5c0793fe5b4ab3da3ef7d264c4cd675fb06bd21ac21d7c77a2ce73"),
		expected:        decodeHexString(c, "0f2c2dbe5da66742e9aaaa48612a0e07dd6914942c52e87d3266150013aac313de1717088e01b93dd0d8c6abd0c5d63d56495140458c4a980ab4ff7d989e00e4"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_5(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "9a1a99f8f2bb88a745bb24cbcd4ac3aabcf9ad963705ddee887ca1792a1fa9f0"),
		nonce:           decodeHexString(c, "a9100654abc3acd6fb0c344404edb974"),
		personalization: decodeHexString(c, "2faad0ea6eb2b5befe024119203bf79c13ed1fedacf5358810f4f870175e08c1"),
		expected:        decodeHexString(c, "2798f2e4b66bca6212e1055d100ffa4bac733b05dca19a3602fcd5e684875a7c58c1e7a8941a7732124daf13f97c57cd7d14e4037abf2f8910281dbfdf9113bb"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_6(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "9f63a3282cd18a09a91a50086ce0e8d91be04261da623875b15f9bac5afe7c6d"),
		nonce:           decodeHexString(c, "ef81c6f367ffbd5d97cddf819a99f996"),
		personalization: decodeHexString(c, "f18fc500cc260beb53aa6a1d1ebd5baf6ec3ad79578a71cdb276c69d675011cc"),
		expected:        decodeHexString(c, "478fb3112c99363b7c6e5d781ccfc2056694dddb70d5e4e457f22cfb154d1d0173c09daa8806ba9c7cefcdd7a715b3b5391be10c3728f3af25ab2a4b8a8a5caa"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_7(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "1015c9c0d06e6ac7acbbab0d2dd1adea756057cc3938ef8785e3d51e5afc018e"),
		nonce:           decodeHexString(c, "951d7861d8409b0a62b18c43bbb65523"),
		personalization: decodeHexString(c, "e76a06024d00be09c6a22b78b74cadbaf368920b1f74a441d510aee02cae721d"),
		expected:        decodeHexString(c, "e1ad222c7f45d82d159309fea15e81bf9f148cacd57b52f5e8fd3faf73d9b832cc803d4f36e3d9afb03cbacbf72d6b297a01d0f00b825f64e676a7018159da2b"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_8(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "cad936acd1ef9685f66bef62a8bf4f919e18024fda5c63c57cedb702df3e94a9"),
		nonce:           decodeHexString(c, "85d3238afa86c654702dfa793e11d3f6"),
		personalization: decodeHexString(c, "d6324b2a3e38946c7df7c6e8d301f6b5e63e55535c8ae0bed677ce0ea86ea499"),
		expected:        decodeHexString(c, "3003468365b5eccdd642d3d430d4c4d4c089c5dfb7dd4adafbdb4631610805e1587d355b56d659c59fa012e490beae5fc767bdbb34b31a5b304c23dc27b4e9e0"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_9(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "648e4e484b26fcb330bf46a9eee418cb46bdb5664dc0b5a875ee8e5845476a56"),
		nonce:           decodeHexString(c, "ba610a74b66472c69f81ad0fbfe8b39d"),
		personalization: decodeHexString(c, "322f584329dfc1f643f359ecfb9833c4aa286667f2ffe6bed9c4c6e912380389"),
		expected:        decodeHexString(c, "7f8498282c32e6b5351215f90e58dd3d5ff5cb7b61c3fd9e2b93dcbdc5ef85ba9eda9fc534565f77b78b420e24cfc95d7d555b743a80b2194508c31e4c30bae3"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_10(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "c52a301c952550a31ab3f3d3c78919cc5509a5b984c0a47a9cae9fe8dcf6ce97"),
		nonce:           decodeHexString(c, "87e28aeff66c45e409f02f99d459f697"),
		personalization: decodeHexString(c, "956e32d3bd82c6379b2b7a7b67fde6583f98d5bc2e22082e793b95e0dcbb4f94"),
		expected:        decodeHexString(c, "6d0f4a8cb3061b6af2f15d4bf0afbaa4624c063fd7ded3903dc90fb359f350b9ec31a19140bcc4d4eab449a6ce4e0abe6b1e819164d39884420aa20c85408c9f"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_11(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "b52b8b3334bff230d033c89ce5f74d21c1ab07f17acec3094f5437e509139d85"),
		nonce:           decodeHexString(c, "617f1f5840a485d042a5a7bc55ca3746"),
		personalization: decodeHexString(c, "a6f40a52f7923037a9391d450952af3b4c8223e9c19e24d2ba11e20d15964d34"),
		expected:        decodeHexString(c, "ced9f3e3fde7509f5c1c46d5a3bfffe67cacffb103fe4facfb45b66e53d594cf4e192844a1bac8187c65cbc9cd843f2383510d8f6576c39d7f37d000b8ce7521"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_12(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "a579374288c8a2de144f460cb0dc2fffbec70eb759c5bf1712a45b5abf63045c"),
		nonce:           decodeHexString(c, "bdb703e646666d1b2163c25fd563029d"),
		personalization: decodeHexString(c, "5128527d9077bd74b44ab7deadeb3f1fc2099ab711ed44a612cbd67c8e946273"),
		expected:        decodeHexString(c, "8fba160867f257b90d76ec680f3e16904a8d1443b02833eaa78a33007137c422669133d7b3de2035830af79400730ded7149546d56a90f204e86477c42fec1a9"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_13(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "1f444891ec1d5d060bd1b037363bd750df9017b27479e21185b9497f6ba0311f"),
		nonce:           decodeHexString(c, "73b5497976d671a260b34fa314ccdd74"),
		personalization: decodeHexString(c, "d70d63e58d23123eba0e5a6bde0779a8866683369ee1c1d1dde81a146717d8bc"),
		expected:        decodeHexString(c, "4356afd04081e8c0f4c446a622cd3f049e7e6e8f906c07600da0671eeae8f0a732f45f7b99ee98c1b061335d0d5896da1e1c081e2a6f25c565fc47dfffebbe98"),
	})
}

func (s *drbgSuite) TestCTR58_AES256_14(c *C) {
	s.testCTR58_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "75d5b782c5345cee69e045844783ba53a315882282740a45b4a04fd9e77ec1b6"),
		nonce:           decodeHexString(c, "175c36b3964726caab9b5a127ee9cd91"),
		personalization: decodeHexString(c, "c7b45b62adabf345dcc018e2eafe3ef212a253edc9460f12a7c0260fe959df6c"),
		expected:        decodeHexString(c, "5c74c046d200cef43613eeb2b0835838609c4ae2697889d51c179a27371b7421d9d28d038bd22aaa7be1e31ca09f3cf3d3ea1d3f18340d9cc73dd8415122caaa"),
	})
}

func (s *drbgSuite) testCTR60_AES256(c *C, data *testData) {
	s.testCTR(c, 32, data)
}

func (s *drbgSuite) TestCTR60_AES256_0(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "4ee68b3352b874e1cc29375028851dee9d5dfd88a40664c79e2b724fb11b2808"),
		nonce:           decodeHexString(c, "1c6a80d82012c39c9f14a808643f08e7"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "7c58d2a5522a88341fb55facefdb6e24840cae283948d53148a384e13b5407d7712c33434bd3d19448b43270c54860bf3495579057c70bff3084dddff08a091d"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_1(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "9442e3f76775093ac2635d9b217974e8c7cc9cce8bba2f04de57432fe6cf0f4a"),
		nonce:           decodeHexString(c, "b94a558de7f887f7f50d3f0cd4f76f43"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "31caaee5d50c6342fd6b3b18d0f88e72b857ed3fe5cbaaf76be1a6acf08551cf3eb15f4b573ca98950c77d30ea1dc3b9fa73335cbaa8e3a5162111269af7333a"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_2(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "27f1cfb937185efff248e1b1188cf1fd9fb489a7c8795ef2c7e0f8a7d7f711e0"),
		nonce:           decodeHexString(c, "e7ac795adcdaae1a93116866c009c5e5"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "14f22ab69b2d5ac916917089827bd657f8d6d1d980b79211bc350b0b3227968d876674297a8987ddb0a944ad1e22df4cf1b612af3f11013a597e2d201ded33c9"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_3(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "6589a90ddac0838c73b7a4529f2c647d707d3f5f17cb76a8df265f264e33c8b9"),
		nonce:           decodeHexString(c, "006a8e6c2facb2355fd6a4638ddb7c91"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "fbed163123f1d040446cafe44b9660a7211d0ff0eeaeba86a612d81d88ee8c6ada33d26115272421e9b84a34d6bd6d7bbbe6043e382f348f0d7daa94dc72a152"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_4(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "e876c10c9e42c75346d593d6eb047a1fb3367398d62316d116a929eb9ececb18"),
		nonce:           decodeHexString(c, "815ab76332db44e713a8e967b20b5c1f"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "7022db947fd418f6c92bf9a12b6d1b2bd217758fa2e36776f35c9d33a489f6913a1d07b4b461a13911469ccf4f3b5211117bdcac052aa8ee0be7e27c4ca2a345"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_5(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "9b3deec23ede86625c5ead5f2c77e605466b602240f556a180ecd2beee201a35"),
		nonce:           decodeHexString(c, "db3d7743defc4dca545fbfc1d37a09c5"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "9dcdf14278bcf9959e2fc07f7b7ee42b9d51114e9665643406036241221c826a595aa1a339406200e23322203e3467de3f14d1e262ed8f4424f3e1e2043b228b"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_6(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "289f5e0e237e45b8f7b22663fbb58ebb9101bfa971cedc1f6977cfef02490953"),
		nonce:           decodeHexString(c, "ecefcaa59d31c969141960f8f5caa857"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "68380623abefe6ff3210e662cb2cd04ef31092acb1dee34fac3a5f70cc9da2c4f93f875cbdc9ef29739c3ba6c900c33750895f107e3bc7f20611c404fb1d3d12"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_7(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "008bf69842d1056881ff2e3571827a6990050f5688377b10b0604412ce86f576"),
		nonce:           decodeHexString(c, "7ac5163c1a96b5f4a5b3d2f2fea95d21"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "18a6261edefd65f634fb8c806c615cd2aa82ca11c608bb6e9e831e1ee336aca2a2474370a461f9fd1a496acbe908ca58eb7beefe2fd5c4802cd527ae0a1cc6a6"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_8(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "419de08d882e1dea8414fd01f1850586e535904ef0de3bdd8b667d16feded63e"),
		nonce:           decodeHexString(c, "5257d3e32cf4cf5a1f6157d60304510b"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "0b7d935414d072ec7b2434a50b9bc90a2987c05fe5596282ff1700f4cc44224779c4ef2f1fe63a4f37ea8891850ba5eaa1f9082ba5b499ba3120368569d09db2"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_9(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "80394569dc1b8daa7d7d42c32f2621fac4d5b6e793e682456314978a2bfb49a4"),
		nonce:           decodeHexString(c, "9b85634262bdaed1c1bbfc669d6cd3e2"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "0821e18a67fd7e78c32f654e7774c1dc1244865e9afc12a430c98c61130f9533e99f641f65458e155bd67fef2f7b35eb7ec7cb38b1e87eb4dfa8a173b23478fb"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_10(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "dcc9adf0029178d4d71dfe39790aa8505ea71681f565294784fcc36312d4d7f4"),
		nonce:           decodeHexString(c, "e4a344daab48eb31d7e8dc38936f1884"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "5eacaed60052ac87a15e42997cfc77b69e898cbc61e36e000af20dc53dd5c8a1fcda924e1030b535c9e4dc87c3ec8d3ce4c061dd46de12f9531067ca87b8bd6d"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_11(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "1fee8b757eabeab660dd95dc0d0d69d182228f241d48c64067f80dd96158cfae"),
		nonce:           decodeHexString(c, "6eb76c5c85c3a35d6f2af4619c54ae21"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "8a6f21156fd61e1a836f6b1f545a06c81008e87c85fe526c3b82dbbbc11943dbaf25fe48ee31fb7018510b2eecdd86b21d1779d5f70ea28b44b1b294a56634b6"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_12(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "b3ed247872210b15c976f423096e3d7f6e626045b2883803262a96aca715e2f7"),
		nonce:           decodeHexString(c, "a1f2b45ca63d4c2a98a380cb24e48bcd"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "be7b5572510a1c1d245cc2725ff927b371eeeea6eacc62ba6aeeb9543af5349fb9da166eeca6e9510316e8d89d5d060234012587ab0baf57a708c1b6703dfc9d"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_13(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "70cc7b7bc803ab30368972cc021c1c51a3f151335add726f5bd0841386c96605"),
		nonce:           decodeHexString(c, "26678777919ac8f1b38f936bc0ab9e11"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "55a481fe7b3c1cd29b942a7164ac7c7413db9711fbab14cf005887cc53fc739c68e37c0b13cb73e53e9d272af803170443ee09110863bf7c10c4102f2aced8ea"),
	})
}

func (s *drbgSuite) TestCTR60_AES256_14(c *C) {
	s.testCTR60_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "437496cc896a20dd6dbf33cd47532f9924e75f26e43a37dbe9f440f5f61136e2"),
		nonce:           decodeHexString(c, "6540d6c586eb24d164e3c5db938382a0"),
		personalization: decodeHexString(c, ""),
		expected:        decodeHexString(c, "4b648f49a148880571bbf717ab213d2d92fa1f2a6983584bc5f3db212e6747c5ee8a6605e1bd14ab9ecdf52c2ac6518821e0d372aea57e79a2617a5d435c20b7"),
	})
}

func (s *drbgSuite) testCTR62_AES256(c *C, data *testData) {
	s.testCTR(c, 32, data)
}

func (s *drbgSuite) TestCTR62_AES256_0(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "5be576ef25e04ff8d61720fdfee9665362da94ce5486f5914f2410e06d09c73e"),
		nonce:           decodeHexString(c, "7b9ccc3e6d5d7b5fb5d4b321e4ff476e"),
		personalization: decodeHexString(c, "ec2941f8684b25dad39f57acea40bd3646e209911d177714ab92cce13afe75e5"),
		expected:        decodeHexString(c, "1f9780ec93e75db864de37b3f9290c609ae3620fab6cbb6c17f944383fafe0f64c233110eacc5b4e5c4107a43a0ffb00a94e00fa8918f11f4c564f04be7126bb"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_1(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "6c35439f34a43cf789b47b4df091f0d2028b9c8c746584ae7ca717f455044377"),
		nonce:           decodeHexString(c, "79d3889692cd2e3ffda028534a12fdf9"),
		personalization: decodeHexString(c, "2eb682598f5ca061f11e6536fc94a3a36f3df2896e2ec9b57740e67c83424b40"),
		expected:        decodeHexString(c, "3464e75b5fcac3799637a40ccda0781bda21722d39b0692b8f567480a98d90a02919553a38830d7ae4d58ef361377fb1a868410568e7f3dae8bec20023eff493"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_2(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "5bbe84ffad62deb1edf2801dbd472d028f45c06fb8334d141f08c5352cbbd272"),
		nonce:           decodeHexString(c, "94f6ce2a287644acc4575a8ba6782658"),
		personalization: decodeHexString(c, "0e3b68da0e167e011d1ec8dd7d8b9afd4b0b6e42806b6000dd79757509e04f39"),
		expected:        decodeHexString(c, "d3a470ea6f5a160963a79531cee9679dd89e05141224883265f214ce17d836a2fcc3e2870d45662d80240ab57e28f83d07e13af582aa7011f969c8e0e732e785"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_3(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "dc881521d70f4d5b34c956831e2c9536d6e026d78629577033be800785aabba6"),
		nonce:           decodeHexString(c, "a85319194e701c558a15a0fcf3f35f49"),
		personalization: decodeHexString(c, "23ef4cf42fec6f4bd15c6afc0e1cae1a47729e2f91019094822a9ea4024bcc35"),
		expected:        decodeHexString(c, "74b03616f6d362a39ff33ab8593b7a87dcd831faa056238a445d0dc623fc47725c2ddea70a5f0b8751a130bb2728bfc3de117677b3886d4d6715f2d8296d71f9"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_4(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "5c3cea9bb9432f23600979d7c4511d8561bf7d88d229f35138ac19646c5c9eae"),
		nonce:           decodeHexString(c, "c098b0aedeb0e7c3a44da29c678f2b19"),
		personalization: decodeHexString(c, "503ef454d59b0c68d23a4bae0715a8963ba7bc70325de3fb59831c907abddeff"),
		expected:        decodeHexString(c, "e4408ca3cc4aac0645ddc3a6612d500c2686299b4fb628730db549fd49b1a8ab8ee7dca0fe5e732edb5252e69b0d9091034248c16548b0317b660b8a8b8e39ba"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_5(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "6ecd2eab50570cd5320cedc36ce4ff7ee311b9ddca8cb071eb82ba04e2737bce"),
		nonce:           decodeHexString(c, "7c5563332425065ebdcd721faf8c3ff5"),
		personalization: decodeHexString(c, "f51ddf774db14f406392bb6c6ec53f231ba4c3c418ee054d76a32aefb71f2b1f"),
		expected:        decodeHexString(c, "a330bb0b8ab6054e99a4404ddd862b8f24039141eb89b0f6577eec2535c1b98b964c5d6a827f5613607863a59404cea55db79be1d08c666b02f5c2cc6a0f7f84"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_6(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "de5a1875c575f7d38ec13ff4774c35b8c6b891b0bdd36042911d157bdfe00c53"),
		nonce:           decodeHexString(c, "f8d1ad3c1510f8e5d0a8f05f439924ff"),
		personalization: decodeHexString(c, "1d320575e79b09f546203bc5d5b86ec0f762675356d84d6e7f57e57b77cd832a"),
		expected:        decodeHexString(c, "1e749f2085097ab5462a100e8a86bd946d29232162280593e900ac778429a3b58932a51189b17238927aabb694ec481dd66c3d937b600fb626e02181d0020eb0"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_7(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "60152697eaee41d5f572edc08dea0ee31f4943446f5573fb23f5bf4ae98825c3"),
		nonce:           decodeHexString(c, "0213d591a22bb0baef78fa141dde5b72"),
		personalization: decodeHexString(c, "36596e0a5529b22b64d08763564676c6f429861e93ba28a43ef7cbb04b0349bd"),
		expected:        decodeHexString(c, "b33f03996ace84097fc3058c52e496c4b3b1337bea78e0c774bd15b94c2704a5d2017ba8920a915e215a757fbfab2d8fad6259481b4f113a8a62e2742d8623eb"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_8(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "e67e0aff1453e6714444171822e66ef7c62e84ca66330d48d838403bbeb4f5ed"),
		nonce:           decodeHexString(c, "2bfd344e1cfdb12b152b84e9a8904efe"),
		personalization: decodeHexString(c, "9e9a6d164bae83b080e46b23dd1e531fd680395429f2e760bd970ff8a2f8c7ae"),
		expected:        decodeHexString(c, "9e244cdcb0ddbbbc24858a52c78cc069c344e465f533980770fafff89f22aa52656a4ff2046bfd3cc7d0944b9543ac0e820f0e1920d54b9218f1fccbf0b271ab"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_9(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "9828238a3a1a6704bb296d121006724f582e5b70ef7531221d6ff63411e60cd9"),
		nonce:           decodeHexString(c, "6b5456cc71b79b41d3b26da1320c6531"),
		personalization: decodeHexString(c, "e5e64fa551427db8f818ce5bdb534ebc7055b5981dc36f6496dac9b89bea1e61"),
		expected:        decodeHexString(c, "5a69b4fe77156f878235ece127ad6a85b0e1119d17a08b21977615d8a9bb8ad7caf35d899dd9aa210c57e1104f9a9ecb7310ea6d54c09884f941c6cbff3da5a0"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_10(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "0bf40a75a2ae58035bbff7664fe2c958adc57900213af9cf4cdcb167b557bbaa"),
		nonce:           decodeHexString(c, "3856fc02dfdefc9c6d003b1fef6b7197"),
		personalization: decodeHexString(c, "b5936c72db21716cac94444ca4a6fdc9f550d0fc7131ddf2709a7fb544d9e2b7"),
		expected:        decodeHexString(c, "f6d9e5a5d1d87f33655c9cba1e936304277d6998488dbf22420c73c82b0744c9cbf21ff82e799d29c1d6e794451ee7e53640f550bef721f0394d73a45715ebfa"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_11(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "98c6b29f701fe30c41d34d7f1da36fe2a719de06e874ae80684399a673a97d1f"),
		nonce:           decodeHexString(c, "0369f7b49f40367116faaa6dfe4ea3f3"),
		personalization: decodeHexString(c, "0e48ab3e507d5c74ca5885a6848c05f7e03bc2a0fc4792aaeaca2a1699f5d64b"),
		expected:        decodeHexString(c, "5f93adee1105aa589a3c54dc0301257821b09367918760b75b8379e3f4d30469c497b7dead634a6eb75e544d7aab18eb9dd2f027c74dc2b447f9dc98ace5bba2"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_12(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "b871d4b9c785a5590e5540d41e574414779760d93fad611c8b40bec74a82d3e5"),
		nonce:           decodeHexString(c, "8b4035fdaf566ec2d5aeb05671dee042"),
		personalization: decodeHexString(c, "8d9adab4c1eb1818fa7cce9e17abb4a2b4b57209dfc251324244a3bc2583a361"),
		expected:        decodeHexString(c, "4f46f9eea62208089fc87d5c7ada6adfd17fdb6192e02a1b90716602c6397968181561c593757b4bc5869647ed7a2ded4f217445ee5998f7e9cee892fface9e5"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_13(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "591be0b5fdf6d1b3b7f5b133d9e05f81ed45f252ab1342a9a225b4312706a2a7"),
		nonce:           decodeHexString(c, "1998ea95f016e68bec009d58aaa917a9"),
		personalization: decodeHexString(c, "bb41dcbd6ba802255a85d4cdd9b38851d636ddd0a83e65262e3d2f4c7dfca497"),
		expected:        decodeHexString(c, "0dd6ba8ba9526222c436746630f07912d54dcd9c0cd4a2e95f0276ce07e031f461b35af87d0f694ddcfaca8fb65cdd657f524f3cb1d92858ab58767acfc2d23f"),
	})
}

func (s *drbgSuite) TestCTR62_AES256_14(c *C) {
	s.testCTR62_AES256(c, &testData{
		entropyInput:    decodeHexString(c, "622a613173df7fb900894609e24440c3f456a3eaf9e38ca988cd24035e3927ee"),
		nonce:           decodeHexString(c, "377c3bdeb56a6fcf9382a242b1036cc0"),
		personalization: decodeHexString(c, "7aea3ff421447e01333c14bb5fd652b159d1f5324ef8d23e5e36081d848caaf6"),
		expected:        decodeHexString(c, "22932674bafdbd4f7a343c5205aaf095edfcf712ce010c11819cd9171d1133c889220a1f939cad9f098c504e62f98a638827d736fb9960829b3826fa046f31a3"),
	})
}

func (s *drbgSuite) testCTRAfterReseed16_AES128(c *C, data *testData) {
	s.testCTRAfterReseed(c, 16, data)
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_0(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "0f65da13dca407999d4773c2b4a11d85"),
		nonce:              decodeHexString(c, "5209e5b4ed82a234"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "1dea0a12c52bf64339dd291c80d8ca89"),
		expected:           decodeHexString(c, "2859cc468a76b08661ffd23b28547ffd0997ad526a0f51261b99ed3a37bd407bf418dbe6c6c3e26ed0ddefcb7474d899bd99f3655427519fc5b4057bcaf306d4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_1(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "1ff8f4a85dbf2f6bb2648967419bb270"),
		nonce:              decodeHexString(c, "b0cdf7bc47ca5f8b"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "f90699441c1ece41cf1f6a32e4948656"),
		expected:           decodeHexString(c, "d9ae8b33f1a10cbf516d97b9ad7baf0d596a081a0ff0f4717674239b9e339354d813b2bb71c10f7d2e34994e0030e4fbfba6438d077c361745993b9d6f669b24"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_2(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "7a3b24c17b87513675c431519e771ce6"),
		nonce:              decodeHexString(c, "abe47800414d25dd"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "b6ffefc408e41f77e2cad479a669274d"),
		expected:           decodeHexString(c, "cdc469c1547903b9fee583409d411e0ac763a00cd687d4f8c811e9c74dc3b78b27b66fe66a249b4178bd3bd08008ea258c5a908d2ea737158d163d1f34f93ea3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_3(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "102e3428b25fda7529248b67bd1bdc93"),
		nonce:              decodeHexString(c, "62d6aa84cf51ad73"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "00d7af7d1f9e36279d07034427a9b5d0"),
		expected:           decodeHexString(c, "6d7ba725c81fb0c8de32f82884185b9eb273bc1ec13e4ca1a0370594cf9c5bd6c27c371826ad86bb8f5d78c697aa9e2927e1c9b2304af0288efe629eff137c43"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_4(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "0887ab2c94dc3bf43ffe8a1badbd5bf2"),
		nonce:              decodeHexString(c, "511853d97431057d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "1733e3828a48fd80adecaa903823ac5b"),
		expected:           decodeHexString(c, "79df1b0832dedde945ed593ce7ebe279ca28a37386d4c20810eded1237cab0b46f1b6b2a212e91ab2150ae77cbc0aaf7bbbc3e6a776a44eae1ef30165a3b6c41"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_5(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "a28a05fdc64da83ecf4d11ffb6173645"),
		nonce:              decodeHexString(c, "2a89e8a1cb2691eb"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "9b82f296bf1defa8642cbe5fc4c7c868"),
		expected:           decodeHexString(c, "606eeacb0e6d8f8e0f6741cf427a935aa6c789e8deeb8450ca8f66bca86e924f42fc3c61c7f69004e247b676f55f4906d2158cea352c607f7fa0299ed09e6f05"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_6(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "71f005df033adf9e1195911611ca51c9"),
		nonce:              decodeHexString(c, "9a33209fdfc41a3c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "88d1ada2990dbfb6916a9ece78785689"),
		expected:           decodeHexString(c, "a24982369ee65868646c8e0c279dff3011605d3399595c455fe20ed338f8c2a51061aa7b40f4e07e86378e896623d46f85aacfa6af3a1d66e22e7cb561e135d4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_7(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d5efee24d6b81f4975393f7b82603247"),
		nonce:              decodeHexString(c, "b3ab7c0e7c280909"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "7d8cdc5f4fe607ef6bc1e2938c27f0f8"),
		expected:           decodeHexString(c, "335e1385201962e636116d6d5ff5fb91ed268388a26aea34cfd354895d9b76a514f34cbe9d3c863178156bb6fed7368a94b037101b26db0d7e6a8a009256db35"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_8(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "006fe4c4c4c016a51a0eb0f35495239f"),
		nonce:              decodeHexString(c, "990d275f3c27c691"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "dd4b17e7d3004c5324ab0b98b75b04fb"),
		expected:           decodeHexString(c, "513f6996d956eeadb5c330c342058fcc465793b636fab74b1bb56e496cd25e5afc7bfb3cd304bb817bdafe7e8d666260f813ec5701b6b4e1b7f872aeee09f363"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_9(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "3b0e7380240061e38f6cb81cbee4bcb4"),
		nonce:              decodeHexString(c, "e882e80e0a2fb7cf"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "b18597904c54a88267cb2651bd5a9203"),
		expected:           decodeHexString(c, "563b0b8dcc828e4293563ff6be314d0ca8fffd176b4ada8d270040863be632540567556702815a0aab40cdc923eba1f5a1c9928494d4011b63f17372c5bd48e7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_10(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "60f5fe65d5cefb731b28179c35b2aa3d"),
		nonce:              decodeHexString(c, "72919f9e097e7385"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "f35412d4ddfde38dfa3f61aa8f6eb805"),
		expected:           decodeHexString(c, "eb7aecb6bcd36dd2fdd32472a69c02f9e71aff13a3e047b58769522d984b608d334e513920f8713d63143c971875b781f3584bc7de352add787820c5f8aac029"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_11(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "52926e414f4c172b8c490ef13f345f08"),
		nonce:              decodeHexString(c, "41d1fa4b9b060f3e"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "fa9152731e26d374c52041abe178cb26"),
		expected:           decodeHexString(c, "38b94e8cb5cd2e49070af89cf208b99f40b0a780578f11465219c7c9b293609270792e6d9f457e629db372dd0f7cee2300e7b97a35ae47a2f75cf922ee5e2906"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_12(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "a76d0c66be9abb6081b52fa5bbba7600"),
		nonce:              decodeHexString(c, "340f89b76e467bb1"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "beaafe74dda0b4e7cd9a24aeb82a15ea"),
		expected:           decodeHexString(c, "0f06f4356ec30e462b1dad3347a929a7d7fe931a377f09e3c3f6d281181c079d2520f78df1d50c6bb73a29dbba58e16a394f8bc7f1dcad3ba8c460f1443f07ab"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_13(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "08418bb6476455948cb88521d1515dc0"),
		nonce:              decodeHexString(c, "8484164b1afc42f1"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "ae4a02a40636024fce89335f6ee11df2"),
		expected:           decodeHexString(c, "59b345ad05a9d2f5c42adab3e790e1e72d0022a2cf8340cb5cea9382ad97e1ac6317fc02850951fa05b9119d82af59dc241dfb624d8cc437f9a8aee3e1c64e9c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed16_AES128_14(c *C) {
	s.testCTRAfterReseed16_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "c9b8d7eb0afa5889e7f9b78a50ed453c"),
		nonce:              decodeHexString(c, "3058ba347ecd11b1"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "643686b86266d9111f29eb389e1184b4"),
		expected:           decodeHexString(c, "0a8ccadc1c5cbd20b8ce32f942505e654b91a4e9410e0ea627c961d632d3be71d6a7dfd64b8f70d28ff91869b92ced908b454936b6d18fcddd7fb77216ccc404"),
	})
}

func (s *drbgSuite) testCTRAfterReseed18_AES128(c *C, data *testData) {
	s.testCTRAfterReseed(c, 16, data)
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_0(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "070d59639873a5452738227b7685d1a9"),
		nonce:              decodeHexString(c, "74181f3c22f64920"),
		personalization:    decodeHexString(c, "4e6179d4c272a14cf13df65ea3a6e50f"),
		entropyInputReseed: decodeHexString(c, "4a47c2f38516b46f002e71daed169b5c"),
		expected:           decodeHexString(c, "31c99109f8c510133cd396f9bc2c12c07cc1615fa30999afd7f236fd401a8bf23338ee1d035f83b7a253dcee18fca7f2ee96c6c2cd0cff02767069aa69d13be8"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_1(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "c6c0853c725e9dee29e9431e8deb2c99"),
		nonce:              decodeHexString(c, "e2be4b86f62a856b"),
		personalization:    decodeHexString(c, "5fb725462a228c2026c1b8894dd3406d"),
		entropyInputReseed: decodeHexString(c, "b8b4b44324a8120c04d6eaffd017fda9"),
		expected:           decodeHexString(c, "e3e71a8dfe7999570c94f7842a0ff975f0e03c32529d69444c246e82458c874a951f2fe6eba366714609473d599e75b5cd9693f769692144652cc8fa39368300"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_2(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "b5186dbe9a099a85f5a353437169389b"),
		nonce:              decodeHexString(c, "c65810e4bead9f2a"),
		personalization:    decodeHexString(c, "4b0cc2ed0899fac025370a796fe824e2"),
		entropyInputReseed: decodeHexString(c, "b9e703daa6ba64e33c6868009e65d392"),
		expected:           decodeHexString(c, "05e452e3b8bc0b60014f38c324cc02c543a00caf4944b8a381559c9d50d9c2f342f545adf87d8622ed18f70642a6dbe3b5c7a07aa070ecf46acea81353f74de9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_3(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "c2f0cf9a5eae59b0cf9a2e72b3659f20"),
		nonce:              decodeHexString(c, "8576f1f1744e3a02"),
		personalization:    decodeHexString(c, "7ab4469c271333629c51a7c2adce9132"),
		entropyInputReseed: decodeHexString(c, "1cbbe2e15d42b4ea086e701e25b54152"),
		expected:           decodeHexString(c, "ba010de5f90cfeee93d5877773ca60cb3bed7cdc86a6b1ed67bfe2a09c9280a9b1635d6f6005d02b22b22026182d4cdb8607d1e85b92c5f3c56385f9ce71e649"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_4(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "f29f113e281c1f9dd4a907a58c32f238"),
		nonce:              decodeHexString(c, "cee12afbd16bfbbf"),
		personalization:    decodeHexString(c, "10376776992d0aef0587b3684a9a8a0d"),
		entropyInputReseed: decodeHexString(c, "cdc7bec4329099f78bb1c3f4de178c33"),
		expected:           decodeHexString(c, "f9ee43430d028de1d495d61f28260a49d52c64a12f52dd384687ef57f8b98dbf9576710ee6033fa5255b0d038b226a36166a0c278aa3a432aa2811d3e118bea6"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_5(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "59e8c9a8ca4e44f30db47bd093274109"),
		nonce:              decodeHexString(c, "c36add93b496f65e"),
		personalization:    decodeHexString(c, "15a19bd596cd9c1340b66919a14ad9a0"),
		entropyInputReseed: decodeHexString(c, "9a4606bbf76f92c4ddc720ee6e4ac100"),
		expected:           decodeHexString(c, "05124320c6dfdb803441e79cc97e85f3066b49c04542ebd3fd7bb105319b244ced549f0d69ca46bb3a7bef33854521dd8e829ccb6320bfb10d0b3ec2a18815b3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_6(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "60fdf3424d6ce229c8773ebea2a72fb0"),
		nonce:              decodeHexString(c, "0b6280f678bbfeaa"),
		personalization:    decodeHexString(c, "f39e3476355cae160aaf2968287eb938"),
		entropyInputReseed: decodeHexString(c, "415127ca26bda6e2c288a4c0f515fb16"),
		expected:           decodeHexString(c, "9c07ccd7d39fe9e5f21987ba482c2d97d7798118f5b5936ad1d987044aa79ed4f3683429d62a2cffb80cdba7e0adb800ed6f8e9923bb4b918a1eb5c92368c151"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_7(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "a49beb27f8e233332186922f3b20eaac"),
		nonce:              decodeHexString(c, "bea24f9a5012692a"),
		personalization:    decodeHexString(c, "115045ea5b7a6f7ff6894188298ed279"),
		entropyInputReseed: decodeHexString(c, "d273aca1b21e7927eecdc023f0402a3e"),
		expected:           decodeHexString(c, "39b23485d826c7f5ced4231ef9b1a1e86c3de191c4c249283e6120bf5524a4e426153da76e098807da0d75339d3c020c4f14842e60490c6f44e9283df6bf28b8"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_8(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "1b2fc3eec116b5559964d6d2653d66c6"),
		nonce:              decodeHexString(c, "6e04941ddef56e69"),
		personalization:    decodeHexString(c, "68be57974c6c74c1c5be0da1c6360e30"),
		entropyInputReseed: decodeHexString(c, "e493c7d991d0430f93c20ae0fb15a166"),
		expected:           decodeHexString(c, "27949a3dc05aba02a624ae002c45a79a65cf6040acca61a0bd5f2ef06a9cf6f824a066656c1b6aac659ee7a9d3308050cc8763d91e48d19dc1df20af14926523"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_9(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "af51c4b1b4bd2119e0b49311279bbec0"),
		nonce:              decodeHexString(c, "4194f00bb0cd83c7"),
		personalization:    decodeHexString(c, "fd052bfd73fe712f4a069d2bc1dc4bc7"),
		entropyInputReseed: decodeHexString(c, "c140c4ce45bf3bfd0a473e5754355b0a"),
		expected:           decodeHexString(c, "6858d2634214a46ec077f86e667b6f3c482c797367ae01d50fc0183887ff9bac3293ce67643be146b86dfb093a5718ab8c12f713ef0cea0d04958ce7a8d8777b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_10(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "fb70a8065eae006270179098a6b67eea"),
		nonce:              decodeHexString(c, "bc0001c0de387f1d"),
		personalization:    decodeHexString(c, "804f5f90bf5a4d7a99f850b882bc7001"),
		entropyInputReseed: decodeHexString(c, "abc77b6f10d0a89452e063fe97f4000a"),
		expected:           decodeHexString(c, "b9f6f51d0de5ffa1fd42ef9a7ad2e60201dc3e4b15536364c1bbea2fcd7c35a77fb824bddc67a690ac51dc2e34cb7efe9f61d5d6d009d961c661de44b98a5e98"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_11(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "92847e86e656eb7f90de2666ba07d9fd"),
		nonce:              decodeHexString(c, "b71f251d206ccf6c"),
		personalization:    decodeHexString(c, "74ba9724a8e8c9e998ca496d3592bd99"),
		entropyInputReseed: decodeHexString(c, "aeed260b730d8f32aaeb0d9be797835a"),
		expected:           decodeHexString(c, "f6a25d68eaeebf57d2d870c81781905ccc69ef7f84f02c7ffaf9868c08a0e443fff019c7b7b649fc50f0e4a95c25ebcd86f94c8403c8a786197094626a98a8a3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_12(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "297efb111809939edebcdf9371cf32d2"),
		nonce:              decodeHexString(c, "b33f34053db7c9f0"),
		personalization:    decodeHexString(c, "d3799c84d31590fbd0d3c4ece3680e33"),
		entropyInputReseed: decodeHexString(c, "c5fe86ee68e459ca3b06e445ea5e0b2f"),
		expected:           decodeHexString(c, "b1835eef69d5406c3648450307ae27d392bdd44a0fa7a70af6c302ab1c3c3cc7c681e190767a38c3ff6b687f3ee12f8555a9d2b4696585f1c811d1329430a51b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_13(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "ed8b28fc95036e5e618695e490b7256c"),
		nonce:              decodeHexString(c, "b995e7c1a88d09ce"),
		personalization:    decodeHexString(c, "a43e3708744c8a7c5feabe66c25c81f1"),
		entropyInputReseed: decodeHexString(c, "a2a59898f453e63fc646af134853e9ac"),
		expected:           decodeHexString(c, "ffff0f1591aa104ba8c552da01c870c51f7dc2a1cc22300508cb4f6310194575099aa4abe7f2918731667394da8b79ae147caf389f4396b5baa88f1917b616f6"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed18_AES128_14(c *C) {
	s.testCTRAfterReseed18_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "7430a658608213045e4ccc24a0568ef3"),
		nonce:              decodeHexString(c, "3fd77ad5b36f2554"),
		personalization:    decodeHexString(c, "c1ad9031efb0fc153615433fbba1087f"),
		entropyInputReseed: decodeHexString(c, "aeca96700a3d66467b2561a3077e6f72"),
		expected:           decodeHexString(c, "c4b248c18fa838acecdf9cc33635009c19885b86b7d1f65c06dab906304da69f4370ac8fc83e4b4721c07351a218afcdaf30d4706324589bf918be20b4dbbe3a"),
	})
}

func (s *drbgSuite) testCTRAfterReseed20_AES128(c *C, data *testData) {
	s.testCTRAfterReseed(c, 16, data)
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_0(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "8fb9573a546253cdbf6215a1805a4138"),
		nonce:              decodeHexString(c, "7c2ce65402bca683"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "bc5ad89ae18c491f90a2ae9e7e2cf99d"),
		expected:           decodeHexString(c, "076282e80e65d7701a35b3446368b616f8d96223b9b5116423a3a232c72ceabf4accc40ac619d6aa68aedb8b2670b807cce99fc21b8fa516ef75b68fc06c87c7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_1(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "a356f39ace4859b1e1994940228ea4eb"),
		nonce:              decodeHexString(c, "ff33e95139f767f1"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "668f0fe2d8a9a92920fcb9f355d6c34c"),
		expected:           decodeHexString(c, "a10661657b980facce7791de7f6fe61e8815e5e24cceb8a663f2e82f5bfb1692062af3a85905e05a929a0765c741293a4b1d153e02147bdd745ebd70074d6c08"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_2(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "9d44a610ec1751cb169a5d83208ccee6"),
		nonce:              decodeHexString(c, "7a507158c989c323"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "d80a20897c167806f2e29199a27ce7c5"),
		expected:           decodeHexString(c, "08e618adca1134cb45a73899865df3ab33eb31a751b4298f98cd1fc8d2cb7c3ca19aacd62b375567d6ba0ce1c08f5fa6e68ffbe76b8eaecdf683dc34e7f8d6d5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_3(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d0730d39d5b5a545d81562fafaa12ba0"),
		nonce:              decodeHexString(c, "6837ef6800d0e74e"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "9bbe6533974a7b1a6fab4ead39358abe"),
		expected:           decodeHexString(c, "f442e8c6fa95343c33f79e32b972e40691e237a25d5526f57724e71841725cba94a5d45b433e80d2f281ae5ab2455ba7779cf6f33700f4a13e4ddf7743d9fe7b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_4(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "bfa911f91952b025aec0d1a56ab1d07c"),
		nonce:              decodeHexString(c, "a50439c803ddf040"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "42920c5e7c08154fdd80c52631043885"),
		expected:           decodeHexString(c, "8c1d408d7a3d0d305c134fda0bf83cb82f00cd7c36a0c9b616ebb23baa6254467f95c84cdb77c596a754cc3bd1047af78f0ee064fb04b5e0f30b4a5e456901bd"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_5(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d0f79a4acb7676071ef771b595016304"),
		nonce:              decodeHexString(c, "5bc0fddf69a10390"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "279eacfcffbb3fdbaf4b9eb5a66f5cb2"),
		expected:           decodeHexString(c, "fe5d3a28e51fb8a8abd0d9c958f18cdb4989e4d5aeab6c79edb5a43e7de909aac157d9094c18c5e5dd043baf065ea2007332ec2e62086b0a61ee49a25f01a238"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_6(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "b5a22a2203644bfb3f2d700530e3f654"),
		nonce:              decodeHexString(c, "06747e4267cabcb3"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "e8a0f995241372d7a306d68ca154de3f"),
		expected:           decodeHexString(c, "c456b404915e2d9cf907c3cc15d5472dd3f5592e8b1555703009e2c9c311839ecdc57fcfc725561c87f6748ff41cd7cd4058c791b7b3cc5d7bdef9b525816263"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_7(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "5046484448d20097de1a5bc3128a92ae"),
		nonce:              decodeHexString(c, "3abc3ddbe20eceea"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "1afbbf9899271cfc4d0d5e1a5926dd6e"),
		expected:           decodeHexString(c, "8c4afb3764125f6862b16219ab4d971c8d11470b45ab316a0c02a177007816a444ec2585e9466bedc27b69647a4c9ddc00d7142e53824f94c8599a8fe0493fd8"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_8(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "ca137dcc6295618b1fd00296b9f51280"),
		nonce:              decodeHexString(c, "e4220ef872a20800"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "509d1a2bbb906a9cc36bd8120783f29b"),
		expected:           decodeHexString(c, "f707b1c7098b2379e754a06849f7e4554531992e976ff1f9d186349f56ac345f3a9b180b4d305213121718369897639b6e96537b557544840a70053102ad8af9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_9(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "085e0e3deea07dd73f59f14e5b7f2ff0"),
		nonce:              decodeHexString(c, "0cd8417be4b4c4dd"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "1cfdfbe27d16e5cf1f40b1338ca88b68"),
		expected:           decodeHexString(c, "41ae72d7cccbc3c5498dfea89539fab8bc4132fe97a4078155a602c48a59afcff8521c9f8b98cce804833df10ed72432ad555332b445159705ca4f083b8c4750"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_10(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "341e1c73a6ea1f2656f5a1fb1d374768"),
		nonce:              decodeHexString(c, "9da8a50cd2ed656d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "060c46770b917e1e52aef57959dd215c"),
		expected:           decodeHexString(c, "cee3336fc4e3d71c6bc38a64dd1c63b7cef02d1f7c7447e62587e55be34b3af11386f5ee74978bafd9771a2cc3ec5cea1ea12ad972b4f4c3219ca06cfda34e71"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_11(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "9703154e413783700f008dd7638b67bb"),
		nonce:              decodeHexString(c, "121e512a318fd57e"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "7ccf5ce243a06f5d8143a771258901dc"),
		expected:           decodeHexString(c, "b2f5ebe6d2a3df80a650813cefd85e139d326813df86b5d24412160c427ead56d754a4cc9c47a238f1bddf234faf22cf0448413c1b0f6ebe76815607eaf50cde"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_12(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "c21f5ba960560548e0e21c5f7ec673d9"),
		nonce:              decodeHexString(c, "f84e7bd5bd4bfbf1"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "267cc99106d5785cdf3071c3dcecf491"),
		expected:           decodeHexString(c, "0a6500cafa9b84701742c0fb627e744a6d4a5c560e6949aa1821321f661e5e1c52a38394aba13b35dcb514289fc8c284b2036dcd8191937c84055723f16a908b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_13(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "284a67bae27ab32bc50162c004cb92ae"),
		nonce:              decodeHexString(c, "04bb326459a887d1"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "79702b24df8df1bd7aba9c32ae596e6e"),
		expected:           decodeHexString(c, "1899d5eed8c2f57e3ed10e780043a85366f1f7c2abc515bcedf66d0b16d698694862adf1a338fb72dafad45e8bc96240d6ac2f3d8bc15374f5a8b48bf6bcfae9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed20_AES128_14(c *C) {
	s.testCTRAfterReseed20_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "dd2c72a4397b13be4553b5c49a2f3beb"),
		nonce:              decodeHexString(c, "bcffaa34aaac3254"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "90874287a689bb9d6853330ef774ba2f"),
		expected:           decodeHexString(c, "a1533421971e2ca2160f2c50c7d609b1b6f56bb7dd7feb0102d9a49952f7aab4a9af4ab5520455ad4bdb242d6888daddb923cf08ff121df114d8da4375f3ef00"),
	})
}

func (s *drbgSuite) testCTRAfterReseed22_AES128(c *C, data *testData) {
	s.testCTRAfterReseed(c, 16, data)
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_0(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "13fa3e445aa961eefcf6016e499f559f"),
		nonce:              decodeHexString(c, "30dc8d2604a56005"),
		personalization:    decodeHexString(c, "da2064c659de89b4f0cf658d4354c280"),
		entropyInputReseed: decodeHexString(c, "0dd9fb5e7a47e28cd49297a6c13d9fa5"),
		expected:           decodeHexString(c, "659e9210052d6c5b5fd5e49c7f6bb534a53e95f31df0eca7b9968e2cf3d5fe7b4d20b69726db5e2c8a80e8b6f60eee71074a9fcd264320b1c533af92c823ac7a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_1(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "99acd51ac4a78edae441b4d1182cd976"),
		nonce:              decodeHexString(c, "fbaffd7749573bd0"),
		personalization:    decodeHexString(c, "0a7f73966dc7f9f41482a870c3a8eace"),
		entropyInputReseed: decodeHexString(c, "725cf181229d41741f02d147d7f0bc1c"),
		expected:           decodeHexString(c, "7fd550147ec824118b644f83e6a0855e4167a1f6496fd6c0d342db4ab136cde96e9abc5e759c7d28b78afd697127946282a32024e3855fc0b57b36aa9d0ed3a1"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_2(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "3af67d00fdf0e73f59591d57b8f6153b"),
		nonce:              decodeHexString(c, "6981c884476120c4"),
		personalization:    decodeHexString(c, "f95276305a1c0e3da2de5c70a7ba970f"),
		entropyInputReseed: decodeHexString(c, "8fe4a3d6a0199b23545df99ca6ef900e"),
		expected:           decodeHexString(c, "69af02a406f263eeeaf7a4c4801bf39c3f440db9bc31b523831717f6948987797b347f12af9b4b4fbeada3a4d416bf986aa9248937c318e93e55039b649024a2"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_3(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "a2136b0d5ca7834940d9069548f5dba5"),
		nonce:              decodeHexString(c, "90b06592e231547a"),
		personalization:    decodeHexString(c, "f9c2dff3992ca9aed00785c216d2ae18"),
		entropyInputReseed: decodeHexString(c, "ddd3757477d8e3256184570df1a6a44b"),
		expected:           decodeHexString(c, "0a727d3017a44771f98442fdf5f1bab06665f93f19d2e1e6ef91738a037b98b83aac19e45bde4791cf74168a260ebb1e058556d31e027db437b828222e515b69"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_4(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "97163af45ad833602f32f938855546a9"),
		nonce:              decodeHexString(c, "37931074ded1f778"),
		personalization:    decodeHexString(c, "386a07fe493f90ee72eb3a652c9558cb"),
		entropyInputReseed: decodeHexString(c, "62516c22a736d6984bd8d3f8d9fac1e7"),
		expected:           decodeHexString(c, "b9053e87e82874155e5deb8fc4499aa81bdf24bac3ef399a292b05d0b879fb75a53f71bb48ca5abc558347d6db37b5534ccd7ef08feb3a0e8a7401a4da3c9c59"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_5(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "e541275a47755d4ef6ab52b81b984c99"),
		nonce:              decodeHexString(c, "5d33837679cd5dbb"),
		personalization:    decodeHexString(c, "135c22e4cbf8137dabbf2cefeeb76db6"),
		entropyInputReseed: decodeHexString(c, "0ec9420c028173b97d1e9615938a6129"),
		expected:           decodeHexString(c, "ed765814cdde167b62e4b0fe6ed8450f14e6f116454570ada2af1394aaa5af68beec0e2f70933c5b973b5e1a7559634318f4e60428f930acda513c20795bf059"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_6(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d35a70ca8221c1868f8d318e52175baa"),
		nonce:              decodeHexString(c, "d7209b944bc4aa3a"),
		personalization:    decodeHexString(c, "6f4e198752d3427a362a1b01fa41950d"),
		entropyInputReseed: decodeHexString(c, "a6849c1caba2f2629858827f92da049c"),
		expected:           decodeHexString(c, "49dc8751e8fea3f3f15cfbf207c377b49a5a9afd453dee375189260a62578d18c4db515791319c453e5b412eb4e369134940fe762aaac8f1316eab8d11309a63"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_7(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "2e0edf2c26ba187d9d409d203b0786ac"),
		nonce:              decodeHexString(c, "dbb2e142a8916db4"),
		personalization:    decodeHexString(c, "325596e2e104f8f6c1054236ebb20ad2"),
		entropyInputReseed: decodeHexString(c, "f61f1f5d07ce98c6e8b0550e07f22dfc"),
		expected:           decodeHexString(c, "babf1b489b3d540995904ac0c93aab228253825f86f79e0fcf530e32b2fe68ebf8cf3a8da30a49afa72bae8b36e35ecf466c677e0e328f574be9140726ef2159"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_8(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "1d4e201ea2a275b4bc62016f902c146d"),
		nonce:              decodeHexString(c, "1347f044da76a495"),
		personalization:    decodeHexString(c, "8c93078f9edb99745442897e7f404c95"),
		entropyInputReseed: decodeHexString(c, "f66cb678e4e333f458f38c84dc8c8bef"),
		expected:           decodeHexString(c, "ddf770d0b2ebb0190d32daae7e1e1677797408d7c826cfeb45bde2f7dbf24473f1c6f35297f498c3e7078d78048d6cda7afe9b6ce6ff8d3f972e8dce7675a78c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_9(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "b05f65560706aec6aa69a53034981a75"),
		nonce:              decodeHexString(c, "978d1396087c7995"),
		personalization:    decodeHexString(c, "7505c5aef50b6556535ff1e79a9285fa"),
		entropyInputReseed: decodeHexString(c, "b24016ccb9b5321749d382f78835adcf"),
		expected:           decodeHexString(c, "eae3632d8b224ca6903f64baf5f75d891640bd38e5714e5cae75f77f0cfc38f06e5eb9c267df6b223933eea5b97377f07a9363bd106872dd86b15ec332c8e920"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_10(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "22a786b99ee8f874311921ceb09b0b62"),
		nonce:              decodeHexString(c, "0212c9fa229e40c7"),
		personalization:    decodeHexString(c, "4f0244e27b0804306c601e84cb919566"),
		entropyInputReseed: decodeHexString(c, "e79e8e72b2cff50023f643deaf2fa325"),
		expected:           decodeHexString(c, "c2049544b5c6395d5ff8e3e41b05ab6cc2727d0d0828fb7f9c46d9f998971f77f48853e307394b2fda282faab50cdc7329a52755203d0b52b7f2a635cf42ca04"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_11(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "ef179fed1a28c91e8976a5fb93d8a0dd"),
		nonce:              decodeHexString(c, "1aed84cec7e6d791"),
		personalization:    decodeHexString(c, "7d9e67a8f04a2203732b4d3f399d2291"),
		entropyInputReseed: decodeHexString(c, "dd3a4947894005ea82fdc3fa1454933b"),
		expected:           decodeHexString(c, "0906cb8844149ebb326ec2998ce2bf1c69474db7da02bc86629e4970eafcfa023019b60afacaea1a2e10bd3fa1107ee02cdf05f1d930a3a8c61a2cc7c63ea5f1"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_12(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "f186bdff36beab82ea451895864b3b35"),
		nonce:              decodeHexString(c, "4449d139cea2458a"),
		personalization:    decodeHexString(c, "4aed4e99555409319e9c525777abbd2a"),
		entropyInputReseed: decodeHexString(c, "67485699e30bf27eb684d6663d0507cc"),
		expected:           decodeHexString(c, "baa141660ead5a6ef41d0acc13e52619a50e262fa71f929bb8e4d9f1ce34e1cabccfed6053c28cbf7589a67696965d0131b223ee83c270800439c6c252fa5d82"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_13(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "0ba7d625a1cfb900e1cbc780968a90f2"),
		nonce:              decodeHexString(c, "f4155df46daee1a8"),
		personalization:    decodeHexString(c, "2dd3b5dfe0ceae7f3821d87ba0308b92"),
		entropyInputReseed: decodeHexString(c, "94a5622bbfda8149393852daceee855b"),
		expected:           decodeHexString(c, "0789357148b9644ca0a5308da2f0c51fcb6a991d09f910b402158059b74125f752895de0b1c70794349b02d72138ad68241cf4c28f8328313c1fd211b1c7c257"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed22_AES128_14(c *C) {
	s.testCTRAfterReseed22_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "e2866d63f6e49e3bd4d94f28e77755f0"),
		nonce:              decodeHexString(c, "f31f809e82177434"),
		personalization:    decodeHexString(c, "05a4ee9d4e41b1d9685ec3a8cdfa542d"),
		entropyInputReseed: decodeHexString(c, "01a4f6e438bd52751fb6dba25e301ae3"),
		expected:           decodeHexString(c, "509fd9086802dfa8a6042e21c588f79606b21e7e2e3985498bd2c23098be88fc8afa2caac0f4004be3d03a2a6abd5e90e3aab7e5797eceaaf0e3835408171d81"),
	})
}

func (s *drbgSuite) testCTRAfterReseed24_AES128(c *C, data *testData) {
	s.testCTRAfterReseed(c, 16, data)
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_0(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "845229e42c4b39c41573544fca838c2f"),
		nonce:              decodeHexString(c, "bcb23a5b7b280c41"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "e549f5815ad32953b115b8da6d2c3dc9"),
		expected:           decodeHexString(c, "45338bb67d2731c9f62ead904a61eb816f93efe0605fb4495f92521d9553adfc5c1b0230216d4c2a384f7ae162ff6319fbdc4e1108446f3372fb6ba5cd6de21e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_1(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "61701c00f47f25493aff86627e13403a"),
		nonce:              decodeHexString(c, "80ae97de15473f70"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "82b57772753df2e7776d63ee975116b4"),
		expected:           decodeHexString(c, "8a5b2099d291bbd21bcb479d37a4082f028dcdf1c042cebd5ec2ddbf5867b5ee60eff1b1dcddd76fe4416f229cab3ad65abc9b57cebb318909feae4c1a2db1c6"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_2(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "b579a25f3d285bf8e5c652c2432f33ae"),
		nonce:              decodeHexString(c, "261291699d891889"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "8a7da7ae82372e50dca2dd7cbf4a97e8"),
		expected:           decodeHexString(c, "8ab1e75d312aa8a1fd7d29f15270afad72a429a7d09cc0f0c33316a73e33b3a615532ff60cec574b4c2f4ec3ee8bdaf9182b2250b70eaa66514952ceca9dfd87"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_3(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "7653fd17aad224506e9d7670f6509d77"),
		nonce:              decodeHexString(c, "b6205411ba313cbb"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "40455e2f1484c89a0f5451ab09a90008"),
		expected:           decodeHexString(c, "39a46870b16fb134a44ee6239b0a1689915b175de51c8acf12ac54d42a90376530d8d99b2774f2da48c6d0f00d63ec721b292ed3f060eab8490d0ef3843572ff"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_4(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d30aa74144d9e93d9d89fa4f02966c98"),
		nonce:              decodeHexString(c, "fd1a5f8b22c1e382"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "5828a64ef309ec9445ffeb4284f249ed"),
		expected:           decodeHexString(c, "a0aeb594e5ffec5261076120a795150d4190236331bf13cd6b994cbc61c3a7006228c8a6090ccc633877e372e1fa469a753504f0ee9d79467359f71c1dffd81e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_5(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "9c9bf9eff8fb95ea6a445406dbf258c6"),
		nonce:              decodeHexString(c, "efa83a3bab1f4c1b"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "c409193b187075c52403e6853f7b866c"),
		expected:           decodeHexString(c, "9caff06f9b61b600e2108b0b94fb8256c136e5d79a17ec3a5091171accebf1216bb529dd4e9f9debcea9ed32223acc55d8db1a55041aee39228251a5496ea482"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_6(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d36d63bf8214c360d9c82877b07160bf"),
		nonce:              decodeHexString(c, "3cc21048a81aba3c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "2ea08431c84ad13527277a317309e5f4"),
		expected:           decodeHexString(c, "9fec4b24a21aaa365619202b158c85797e64129bb5a095c7a35a2ad5bb4c583f8ca029b69f992ce1a1771ec363b016d20f9b012b78696cf2299acb6957c01777"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_7(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "89d2cf4258d78636bfa89dedd66ece30"),
		nonce:              decodeHexString(c, "67ba75ac919561f9"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "aa1f978d93c747933e9f717e101bf8a0"),
		expected:           decodeHexString(c, "f9451e0387be0b518c538422f4aeb80d775561304555c5957ec9c7d52a8b773c7bc19a21a848d4741c18836396b37011fcb454ad87ec066e40a861dab407ef37"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_8(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "b15a9eecbd33ef601fc4946b78b31803"),
		nonce:              decodeHexString(c, "b145c3db01c5dc45"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "64c7acf3473d0c4ab3d9d20e424b8082"),
		expected:           decodeHexString(c, "590ef8df1309d1f2cff27e405c722329f699cb5672053131228060280f448d1fef8091de469aa6c26dce06def99833802b363dea4d7c4b85bf28149be93d6c16"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_9(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "8d357cf4eed0f104e393e850b91048bb"),
		nonce:              decodeHexString(c, "a99f29509bedb6df"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "5489e4838ce24a12c579fbc10fe79724"),
		expected:           decodeHexString(c, "55d7568606d76fcd236e843dee343890405429b056a8b63ff52c089078195b549de6718b362ddf2f6f4c979cdebe2fe793332b78b892dac8edd22c42a5d06880"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_10(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "6201550244c1a42a4c457a7591d25743"),
		nonce:              decodeHexString(c, "e83427a9f48d567e"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "4449410cfc18830ec18512ffbe64e562"),
		expected:           decodeHexString(c, "09a34444998a2c1caced5d9313a3b8e39a13c8ba98a33e96d78a991055acfb384688b5b3e1b87c2c3d0c910fb9a420499384a49070201b76c14a126851b8f43d"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_11(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "5bd54faac8f1b11fe1d50afd36b0589c"),
		nonce:              decodeHexString(c, "e385788e7b67fae2"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "8726550330d852838a999fc989e4a1a6"),
		expected:           decodeHexString(c, "e87bd188f3feefb4d058bd732a0f225689b8eb31cbbb65b3161eb5f8d523b1befc59b005c8820d335a060ce2aefabaf74a6dd994f8134e7805ac65e506570af7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_12(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "4749f6ae9d200ce0ffe5bca2f7007c5b"),
		nonce:              decodeHexString(c, "90db60c760e3002c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "54ad4ca19402fdbd376006021040bdaa"),
		expected:           decodeHexString(c, "379801262fdbda3b30373719ea2241b8e84fb173565163d963502277fd41b62c2f78506d23ee80c7f2e8b96354a6caa69dd694d8049467504617f694894574c2"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_13(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "184845ac4e884213081cb29423b5116e"),
		nonce:              decodeHexString(c, "299207bcbcaa35be"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "5460f71393fe8aacba6aa2b690cac978"),
		expected:           decodeHexString(c, "019efacf060de6898dc234ad02facc8179bfd7e4fb708f1c015192d43d953e590bf3e8a252fbfaed5d9b5d69c99c2343d9c32c71a9010a12d69a8e6c86e5890c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed24_AES128_14(c *C) {
	s.testCTRAfterReseed24_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "3729fb442bc32e132a5b3b6f379dca0c"),
		nonce:              decodeHexString(c, "5bd4f3fb9ef363bf"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "a51080a35e53e80f63678299eb046aab"),
		expected:           decodeHexString(c, "4fd586f74c278c149d1bee6714a5cb9410c205d042740b454b7a841470b7f987f6af8a680e91a3f2bd88116c25d9e5ec8a78f2d12a1289a01574f2f751181cd0"),
	})
}

func (s *drbgSuite) testCTRAfterReseed26_AES128(c *C, data *testData) {
	s.testCTRAfterReseed(c, 16, data)
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_0(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "ae5a2f9726ea953e4ec057c4c96ddd83"),
		nonce:              decodeHexString(c, "7fe815f735253bf2"),
		personalization:    decodeHexString(c, "d4914e8870c298365c8c5df216d259f3"),
		entropyInputReseed: decodeHexString(c, "2cf2fad8c59c50508608555549cd611e"),
		expected:           decodeHexString(c, "f9120eaa71e3d8543333cbd0a83b46ec86a22200878616106c866e13a8cbd646915ad81c7a11aed8396a25f98b324b5352eaffd501fdc9920b5353590eb0409f"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_1(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "95ee917d266131bfd690cede3e20fe6e"),
		nonce:              decodeHexString(c, "3768c29157e4192e"),
		personalization:    decodeHexString(c, "0f89375da6b2f7aa5e8e86c98e092889"),
		entropyInputReseed: decodeHexString(c, "78e59c9b60282fe55108eb896c1627f4"),
		expected:           decodeHexString(c, "21f45f0f87955385b3a604f87f7a770fdc4e4c2d6ccb7cf53487d99c9e11bcfd886f2223451777080fe49e3b8a4059a83c546610107c9aa9a151522c7219634a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_2(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "244b4aacc403b08f4fd80dd655d2639d"),
		nonce:              decodeHexString(c, "41639d0315113fa6"),
		personalization:    decodeHexString(c, "ac587a5c8dd70b8238bb8ac673ef8cbc"),
		entropyInputReseed: decodeHexString(c, "a4bb89bd6be858c9289053a245c47ca1"),
		expected:           decodeHexString(c, "90b59989325f3e7fda225c2189cf7f6cd4a4c43486b42bb7af90f6f2b15f9de54a3fb4e1c84b76249153cb6f6e2aad8dedfd17b90252228559ecc8a7d718dec4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_3(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "6ec82c6c913e33a11756eef84457b766"),
		nonce:              decodeHexString(c, "1774b27cd4c33784"),
		personalization:    decodeHexString(c, "808c58bcfecd0826b295f40108ba20d4"),
		entropyInputReseed: decodeHexString(c, "f8cb5bf5e68ccdc43488f2578bb4a5f0"),
		expected:           decodeHexString(c, "11e8e816053bc066be8e9706d3f937cbf98e50eca805646138b41fa6a6d4a358077fb794a69fdfaddd75c4bbb5e6f71d7c5e8cb088bb25162e25eacaaa1c4d1c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_4(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "e3f025984fa7c878c0d1f97215abe474"),
		nonce:              decodeHexString(c, "33ccc2ed466c1e7b"),
		personalization:    decodeHexString(c, "37307836b15355447a635f28cce03888"),
		entropyInputReseed: decodeHexString(c, "3346d389742920c81ec7c9319f457f85"),
		expected:           decodeHexString(c, "b1c91952eacb72b4f73b960ce4f228d18ed3cb22e4a87072f89cb865adbf4d38e3d8a09d1c1bb3c71d9e63d7cd59c2544186b86eaf79fa165144ad9c3909321a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_5(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "8ff35a1ac080dff97ded221621166a86"),
		nonce:              decodeHexString(c, "f78aad37939c788c"),
		personalization:    decodeHexString(c, "fbc5b61a7c04851f0690ec4cc8702757"),
		entropyInputReseed: decodeHexString(c, "e251fd54e0e66660d5ea10a7edd22029"),
		expected:           decodeHexString(c, "d4d313907b08293c1a03cd3d9a63192f44677d370ce320501b5faf2cd66fa8fa4fd1dd3f340326a99b8b5e5b02a636eeb04fbb0f31f193be24342e6faa3bc233"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_6(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "8dff9853cca315ff7e45a13b1dc20095"),
		nonce:              decodeHexString(c, "7e160780d697255b"),
		personalization:    decodeHexString(c, "628055ef3a084967aa287f626d0d8caa"),
		entropyInputReseed: decodeHexString(c, "8ad8eb28ce1bad08f4a24f3b653e455e"),
		expected:           decodeHexString(c, "d65778084ca505561355b8c3c8cd0833cc94c68038b00f38ee593f0bcad797ff6ed2ddd98f9c65318094b43c25dc59b6bbb0d31f79c044f910dabbd6c1bcc54b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_7(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "118e19a31ccdd14b8449e6b12e1340b0"),
		nonce:              decodeHexString(c, "22187a3abb4477a0"),
		personalization:    decodeHexString(c, "e7b885952c83884ea6a96c1a51ee7ed5"),
		entropyInputReseed: decodeHexString(c, "8db14d02f1f63c385dfd5592aaea1c2e"),
		expected:           decodeHexString(c, "15478cc9c12df70e182008b860e0c4de5d8dfc375bb686c58eb694f4817b4b0c373de53c4901b64d915bbaf7a41a4360aea9cb996ee3d51556fba9b0c3196bd4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_8(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "83f3f7adeb932e3b803f435213a3283e"),
		nonce:              decodeHexString(c, "248d7af9ce6e4167"),
		personalization:    decodeHexString(c, "39fcae726dbe651bc80e56a7ec19fa60"),
		entropyInputReseed: decodeHexString(c, "0bfe69b9abc8a76c063e06d3617e4147"),
		expected:           decodeHexString(c, "7e6f2ee570137abe6c6e8db2556b2d30f87aea4fea5fa063a110e942aec4a2770b571bad44f07d3989e9adff2918bebc04628a88eeca6c6f2cea7cc338eeb670"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_9(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "2de989ec6d0da9ebcc02f724a4476712"),
		nonce:              decodeHexString(c, "f2f3794d18ac2bcd"),
		personalization:    decodeHexString(c, "c50ac17e4a68f7f02d6820653f0a71fd"),
		entropyInputReseed: decodeHexString(c, "350927eec6636305babd0feee243082f"),
		expected:           decodeHexString(c, "dab204ce72afd6cd256a1754ab7ebe424d4bfc6aeeac9b6bbd06e79f21286345457b5fd46281c145eec450225823c7596e18b7bcbb379845a384ab48d8a555b3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_10(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "760fa0515304e524743ca563baf248f7"),
		nonce:              decodeHexString(c, "4686bf0437a286ad"),
		personalization:    decodeHexString(c, "b03922e8923493d2227a1bffc9ee7c09"),
		entropyInputReseed: decodeHexString(c, "70294aaba8042e770e2edd6a02d365aa"),
		expected:           decodeHexString(c, "eec8b2577cb10375f56f8aa8789cc816c0bfd4b928f2b5afddbe8fddb14b0faa1dfdba5188a2862523cf4c44ce1d90e12a1fcd42263b305e2d85d3c40f563b14"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_11(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "e102ad437aa0a7dbd850608377afef1c"),
		nonce:              decodeHexString(c, "0c1214cfc9b3b450"),
		personalization:    decodeHexString(c, "15a79cbbd9ec839efc49a1030e8c20ae"),
		entropyInputReseed: decodeHexString(c, "8464c0fb7335305f54571f2951272115"),
		expected:           decodeHexString(c, "b6b18e26d90e70e0b786ff640840273ee086bf03205905fe1cdb2e9e892f0b2aa2d6e78ecd3cee2d99286160c294dc368ed45e2eb331d03bad58f4b5cc505053"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_12(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "90a9acc6566b8ba0db4ed5ba9b7d19d9"),
		nonce:              decodeHexString(c, "74b6c943761b9568"),
		personalization:    decodeHexString(c, "6bd16bc6b029ab5f76683be712dc3272"),
		entropyInputReseed: decodeHexString(c, "fee3ad6f484584b3163a213f199a4d42"),
		expected:           decodeHexString(c, "a3cc0df53f0d516a279a253e6853b146a4bded0c270ae661e19faacc5a764515d01566882f88de5160f402e2ce0ec46d1d7d50289446de69ac9d889ae10882fa"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_13(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "a9c5fa1fd15dc35acfd25981dd547b56"),
		nonce:              decodeHexString(c, "267404714f70a7cf"),
		personalization:    decodeHexString(c, "610714922290360d4919a9f1f471bbc6"),
		entropyInputReseed: decodeHexString(c, "9834859ecd852a7269546e825c01d611"),
		expected:           decodeHexString(c, "e72eba34d42b8ea693bce7e876c7b338c1a7191b07c9c53e1c0e27b08bf0636519a65f634feac52e760ae7816699babb0c0268b17238d80ce698308e6b360cf9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed26_AES128_14(c *C) {
	s.testCTRAfterReseed26_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "e761778dfcf12b1697956ba1f49617be"),
		nonce:              decodeHexString(c, "8ce252ad4dc8a702"),
		personalization:    decodeHexString(c, "ceae74c9760fb4495d333f84761a8a30"),
		entropyInputReseed: decodeHexString(c, "1a3653bcfa549c1f98f17f5b26fc04b1"),
		expected:           decodeHexString(c, "67ad31843d23d5b1b21f70876351d5252cf0cc93a91a7a7c8f785189f829349be67acbb9dae05d52b214bce4b310197738377a5f050e254920910b4497d607ba"),
	})
}

func (s *drbgSuite) testCTRAfterReseed28_AES128(c *C, data *testData) {
	s.testCTRAfterReseed(c, 16, data)
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_0(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "b25e5c411c47c68711f62a1c95c6094d"),
		nonce:              decodeHexString(c, "f49dca3c2cd31643"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "84806941378019ab90b1e845f1b4afa9"),
		expected:           decodeHexString(c, "25e548218a2d851d596b02ed1c18acea5c5abbf438fb838ef4762f736eb89fc87068e1455c6ac32d162dc32e543cf5dff09b9d3a19d73b0dd25c4e3fbd0b9309"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_1(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "2d8af566a207095f367cc2fd6b553f09"),
		nonce:              decodeHexString(c, "cf043e8ac56af056"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "667ab3c312fe1962de00de6471735813"),
		expected:           decodeHexString(c, "a99d65d52eb2a0bb6215118751fc3fcd3e5951fa1701bc1d58338a14811991fe43175c4f7d8ed49dc6bea22814701e89593957e5543733f9a5d761ebed5df0aa"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_2(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d27e9dc50d02c15c94fb4c89f0b63f3b"),
		nonce:              decodeHexString(c, "03708637e52dbc43"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "29fe203f73d1ec9d8b2db843cf24935c"),
		expected:           decodeHexString(c, "6d0586b9bbf78477d1763a6b13946dd7ff3b64e1a0812653273b1b8924a61037cc824719fa7a68820d6bebf9ead16bd8e9b168ef94fc0e8c10c7bd8eeec05e9e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_3(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "9d35f5dee9dfea543f125e438b9d8085"),
		nonce:              decodeHexString(c, "a071d5cc0816f770"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "720ea884c253ef2350c2b73757a7e589"),
		expected:           decodeHexString(c, "4d571fec7ac322d1e0da52571342d97525f2a5dd063beb0dbba8f1592ca3263f46596fbe15bb2187b4234e1d249cab0bd18a7024acaa6025565699ec88caefca"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_4(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "4f1702fe2a6a2febfbcc28a4a3cca895"),
		nonce:              decodeHexString(c, "a8602c8329e5ed59"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "7eac9189d1723ecddc4d4c9b7da75125"),
		expected:           decodeHexString(c, "179ac5065c446bb0213a0f7a65ea7287f03221aa0a53e3838cebf109a021b520a366391ff67ea7806b42d6fbe329aa4c7e4732d16e324abece0334f9950f59c7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_5(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "5df01e87e11dd227127820dc1d286cbe"),
		nonce:              decodeHexString(c, "7fd3c3b39aa910d4"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "114ec09430b7c0f43e41b65461c28baf"),
		expected:           decodeHexString(c, "9bd45879f0f0b9faacd7c64743fda1d0037442c3c5c581bae00a4fec9ad6a1175a15fdbba56e6237399281f687022e4a1ca0b45320aed268e6111d92609a1b40"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_6(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "60bf7d609d7810b0d486ad95321c3b76"),
		nonce:              decodeHexString(c, "ab8b8129c5b999ff"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "caed654dd87903a784798f70aea3d5c7"),
		expected:           decodeHexString(c, "a442507dc8e23b83b951af0f25c207b73148f4800d45f3948a088ff3d961d21d9375558056cdea1228f2e93fcde8c3a4ce8c8d7f98028c8dcb1b38e795ebd495"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_7(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "33834550dbb58b77f6a11c5043da2c35"),
		nonce:              decodeHexString(c, "0b315724f1d7d88c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "e2b07a290f047961eb10c26e58fe32d3"),
		expected:           decodeHexString(c, "28e70d99b3efd0d9eec93fc533e335e568eed57c21d7fed920e5544667cfe13fbc063dd818f31cc83127b639458f660f7481663adc3bb1fdf9187fef0a7a9df0"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_8(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "3ed59577bcd9a04f0d87f097b06aad2c"),
		nonce:              decodeHexString(c, "6acc6cf7185ba0b5"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "90d04f840172e5f5ccb34a113ec7fe9a"),
		expected:           decodeHexString(c, "e3b8768023875df81b9fcaa4dd55ed533ea8c05b57f02c2bead327830de5293176523aea974bff8352a784321fe308e1b2897b47078d25c24fb8a171639d276f"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_9(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d1177626cc0676737d879299d01a5a24"),
		nonce:              decodeHexString(c, "ef9efb3a6ed4848a"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "19beeb41b5969df1aba65de85474f6c8"),
		expected:           decodeHexString(c, "eefa3b05f8742c2c31f72f362ac94e9c254d398c2c7b4315b6060e733c801fb956ff0648fa87338d5feb1e3304fabec529a337e29462d36c02a2757d0b31d4c6"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_10(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "9287cc2bb1fee82a6d74ccb85349c75d"),
		nonce:              decodeHexString(c, "222d7a6fe2c16c1f"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "f5f601cebc5a7fb73f14ed2908016c7d"),
		expected:           decodeHexString(c, "c54051b9dbb98b812dcc417b6715267a117209d4735877b49e68e96f4fa7167d6e138269176fbb4d5ebe53286dd0c0c3bc6b7fc6d74264458a61cb67418be4b1"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_11(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d8beea04e455b15c268702c787a546bb"),
		nonce:              decodeHexString(c, "d33f603433f097c4"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "3d0e2090dcabf521160962e5ee912132"),
		expected:           decodeHexString(c, "4edaaf5f468382ca0a8c9b51174d97c106a2ed2ec636e9b4d29c2599c0e0d680aaf7bda15b267f8f161824a01cc60d418a8f13e7cde31fbd2a85272cb5577679"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_12(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "f241f0598f9315dcad17e6098a7d5d72"),
		nonce:              decodeHexString(c, "3a79da9ad1749c15"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "c5046a9d087946e3ac1c5412772d262e"),
		expected:           decodeHexString(c, "07d0b56509418fba49a362eb47c3e230b9329304c0a4478949586cb1fe28b1d0480c4f443b04b0a2dd12ced3d721f13e7f1bc0ed074bcc0847ca0efe7fcbc363"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_13(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "8a40d580e53bdbbcfd9f49cb7656e3a2"),
		nonce:              decodeHexString(c, "d4129fc3724c6d41"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "88eeadc32c3487cb074543abf2247643"),
		expected:           decodeHexString(c, "6a02e4c3c814f75f6518fb2a5295a05c5af491b64389f9161179af5436a52a5ecdbc0731dedaf95a24dc56ec634c9745b6cd333a6de9604d18a5187ac7cc9a1d"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed28_AES128_14(c *C) {
	s.testCTRAfterReseed28_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "0756b7331d1c5b8552f104734a4e4f5c"),
		nonce:              decodeHexString(c, "1778bc4130ecfd4f"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "488ed15dce02e3c2e3ee571260fe734c"),
		expected:           decodeHexString(c, "8ea197963472da0430aab9c31b91b5f95924e7c8a34c2625277ddbdff551b837f40896f133f7f6855fbbbcb197d4b222acb3c0d04d1cffc3411947a58b57166b"),
	})
}

func (s *drbgSuite) testCTRAfterReseed30_AES128(c *C, data *testData) {
	s.testCTRAfterReseed(c, 16, data)
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_0(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "81052d3322a260fcf3339b53f6cf5d6d"),
		nonce:              decodeHexString(c, "57303972ab58f064"),
		personalization:    decodeHexString(c, "c3cee8564d8bb2fb8efe6b16aa2616a6"),
		entropyInputReseed: decodeHexString(c, "c9b103f3779aa63e379841b714196cf6"),
		expected:           decodeHexString(c, "d53712482c2b3336403c40aae156203c5859a6b3fae4b553e72b8a6687c0152bbdafe4099439e197e5d60f8e602e5b550aeb7387a4347ecde6e3a04288390f13"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_1(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d712be6df1e21dedd88893031da2009b"),
		nonce:              decodeHexString(c, "fe560d2141846540"),
		personalization:    decodeHexString(c, "49ec80b4f3f85388d2ea7454ac31bfcc"),
		entropyInputReseed: decodeHexString(c, "b5ed3dd2f5e19cbedb9da7b0f140c182"),
		expected:           decodeHexString(c, "0422d60cecea426612044a3e97c4c962975c86f59b278f80abb5b5d20537b39f979394dc3418d004f391297af0fe5a1742166bec5bc013e2e7329ce308d8a0ae"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_2(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "6308f7ccbfe7469b1a9d97274a9403bb"),
		nonce:              decodeHexString(c, "4f63a862808a0b5a"),
		personalization:    decodeHexString(c, "0abf7c0b6ec5423874c17001e1953460"),
		entropyInputReseed: decodeHexString(c, "f2f88715e198c7e3ee46525040d8d7de"),
		expected:           decodeHexString(c, "008e17da7757afbe5c5f2f790882e3104e234381ba5ae3a99b7a450147973f543c8ae0672ee724eaa40973f0b724610fc512981254a5453443bd68253a673783"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_3(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "e89dac4d14abe1783fe79f9cd7466dfa"),
		nonce:              decodeHexString(c, "632c72696a6575cf"),
		personalization:    decodeHexString(c, "d115426c6abcfc94db8a03440fa68069"),
		entropyInputReseed: decodeHexString(c, "14a767a16c2cc35c389928cc9cae2c6e"),
		expected:           decodeHexString(c, "3a504fd1c95b6cdd1057b3aa6ea9afb0dd53475f774df7797d949853fce538e1a3f1bdfd7bc86fe0e4496f210351afd06462c3e0fcb66a5478cc54c854690a69"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_4(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "a85e33ac6b709b7de2ad773f77628b32"),
		nonce:              decodeHexString(c, "7e5b6a38b6660c54"),
		personalization:    decodeHexString(c, "1abf0b867ae7810c52bc8bfed2437059"),
		entropyInputReseed: decodeHexString(c, "65a8d3d0fea01a05168124184deab4f1"),
		expected:           decodeHexString(c, "b87eba4bcf9b914823818be7d13a73d7b2ad39a0bb289559cb1862552688e0d4229a7b1069e2d459fbbe3852d548fb3160260255194152e62a3cdeb7af816ff0"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_5(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "0fb403ce513c9148efb4973139da3c1b"),
		nonce:              decodeHexString(c, "bb37f73df280c7ba"),
		personalization:    decodeHexString(c, "ad712ee96ee96ad82108815d975f0f74"),
		entropyInputReseed: decodeHexString(c, "fbd8717c36c2fa4dfe1d9a690b5d16af"),
		expected:           decodeHexString(c, "bc16da47e64d4801c4c474088aa53564673231dad9da7e76b4ff8e5be653d561d54acfd2a55e0f4cae45333eef0ccc28abd94f5e118efa7b3c518839b7e409e0"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_6(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "2ba9c09fa833a4a198ff0679a02261a1"),
		nonce:              decodeHexString(c, "764e8c4605711d49"),
		personalization:    decodeHexString(c, "465496ddadd10086b7e6c75780418177"),
		entropyInputReseed: decodeHexString(c, "da681a3f603a2ad1eaae2e7efdba5197"),
		expected:           decodeHexString(c, "d63528e280468c7192cf7739dc3bcdf56b3a8560924c7d1848d64ae40a0e5c124b47d2d1cc439a04d753b3b8858fc227b5df11864bfa224e46bf0a6a01ba683e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_7(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "797b612310f4d89755ef63d58951cc90"),
		nonce:              decodeHexString(c, "21b821f937d6e660"),
		personalization:    decodeHexString(c, "9b9277b3685dc03012a42d2f01bb28fb"),
		entropyInputReseed: decodeHexString(c, "60523f6764d42a6f947939ed11511586"),
		expected:           decodeHexString(c, "4b2fa5538397eb49f6b6d76d55e532b578bfc5a1f288d7ba20c8518166975bee53e63fe7df54ccf69e0b3ffa077278f92cc1448000d96d82cb45c055ce6c00eb"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_8(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "d08526c17883939324fa66fde801a3bd"),
		nonce:              decodeHexString(c, "21c28491a8fe20d9"),
		personalization:    decodeHexString(c, "6ef810477d4d9614398946653fd34712"),
		entropyInputReseed: decodeHexString(c, "40645beb71b40f95063d492ed3bed7c1"),
		expected:           decodeHexString(c, "0b2c243161c8386ec996a8693fc307c0c5a094a023c812337f8d476bef8c96c6f716870276b09009494d64df31e66ff6d067fe589c357eec5ac99cc9a81b2d6e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_9(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "b91afee6057bf4c6306964b86fbbbcbe"),
		nonce:              decodeHexString(c, "090642799a480be1"),
		personalization:    decodeHexString(c, "fcf63a6b30c2fdd9a53555779605984c"),
		entropyInputReseed: decodeHexString(c, "1b47800f5b6189116c341e9b99c43b01"),
		expected:           decodeHexString(c, "9d85b2aaad31e02479e2af3a59cc61da98cdc24650535618318d93dafc18590d5af6a420253d74d52991b1228376198b1cab0e12184385e7c028923c93949310"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_10(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "a9aafd15d109d30cf889e6545792b869"),
		nonce:              decodeHexString(c, "4dabcc8c293b8816"),
		personalization:    decodeHexString(c, "a0e566637074c9407c2f18f57f8b0c3e"),
		entropyInputReseed: decodeHexString(c, "d18b43413bc0c9bdc8b8ca83f22b1c7b"),
		expected:           decodeHexString(c, "4058a7d03309c5494ba68269b487ac6eed2a8f203db91da9e094137e53cd48453d50e3a641dbc19a4a7e6409f4c1f709719ef835cecdb7f1e5c4a3e44f125e22"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_11(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "8bd67614996c00fae1c151fae1b78bbf"),
		nonce:              decodeHexString(c, "32ba6e0150a08c02"),
		personalization:    decodeHexString(c, "5a4328ea5b5cf8e16ae21e8df06df940"),
		entropyInputReseed: decodeHexString(c, "cdd5e249a77709e02d34ae6a4c713876"),
		expected:           decodeHexString(c, "616cfe6127eddccaddaaf816248293a71fed14661ac067ebcbb8b315c773bd54e5416cb1f2f4e3f15e785b9a48b5a310cf241d53b211bc2070c74f53f3163109"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_12(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "123c4344436cc0990b3ade31830a974b"),
		nonce:              decodeHexString(c, "b5cd6c8bc420258b"),
		personalization:    decodeHexString(c, "f9d9b8d8d1d3ef2c3cec38968c406c03"),
		entropyInputReseed: decodeHexString(c, "94906024e8b4f573208870c404267d30"),
		expected:           decodeHexString(c, "3511d41acf56d637a813803d1cdbedfad5a2c322ea64498729014468ee2e7a2633fc37f3ca11be7e4987f530f1f9a96649829eb6da524cf9bd0c7f20b2199415"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_13(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "0cfc00d772f38e58a6efdedef1fac3d9"),
		nonce:              decodeHexString(c, "46996dcea6ed652f"),
		personalization:    decodeHexString(c, "68e39ae143bd7af759339df9cf861122"),
		entropyInputReseed: decodeHexString(c, "5c7ee1192cb3b43d3029ee6ce7859534"),
		expected:           decodeHexString(c, "6fb8871ff319f63119081830d38d43f37740c3b862d6b2a0780d89bcb38921a03496fe8a09792f4f64bd0ca1b6071ffc8db76de1cdd0b09c2a0ea1922ead6043"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed30_AES128_14(c *C) {
	s.testCTRAfterReseed30_AES128(c, &testData{
		entropyInput:       decodeHexString(c, "813ce8c15be0eef6ce6e119a5b3fa23c"),
		nonce:              decodeHexString(c, "e471366732803f51"),
		personalization:    decodeHexString(c, "a35b7d7037f4e1a2bfe3ff17c2cee87a"),
		entropyInputReseed: decodeHexString(c, "8f8daef25b0c78ae5add50cea3054720"),
		expected:           decodeHexString(c, "0170651e7eec1f0011c957313a9c332021964abeba2e626efc59d8b03022f8b2772ce354e5dd7d6c20951cb6341ce9861b35346e9d166d92ce11fbc0ad21b464"),
	})
}

func (s *drbgSuite) testCTRAfterReseed32_AES192(c *C, data *testData) {
	s.testCTRAfterReseed(c, 24, data)
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_0(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "b11d8b104a7ced9b9f37e5d92ad3dfcbb817552b1ae88f6a"),
		nonce:              decodeHexString(c, "017510f270c66586a51313eadc32b07e"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "6d14cfb36f30c9c1a1ba0e0a32c2f99d1b47f219a3a8ac14"),
		expected:           decodeHexString(c, "53fbba563ae014ebc080767aab8452a9f36ce40bbf68f1a12dc0a6388c870c8dfa4250526cbc8c983fee6449903c6bd7c2c02e327680a66b464267edbc4e6797"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_1(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "880a7dee52ba3583044254bdcae89db6f8451e631937fecb"),
		nonce:              decodeHexString(c, "82999b1f6683a30ddf527659df473e2f"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "1153cfd0906d9d82f2b673ccdd928eda89870576c9c478af"),
		expected:           decodeHexString(c, "a07979b2aeb6c48195b917ebb4f3ea4d248e9624a910012900ebbdff1af4a975400dcc510a0f566b024a2382062808cd5b4439dcceaf7ce7e3837ddfc25cb166"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_2(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "b85ea76739baf83c2ce44c9cb3b0d655ef369c63148505d2"),
		nonce:              decodeHexString(c, "dcf5f66c8dc3d96e117ca8b907565879"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "bd67bd3fabafd298d8d2a39e7c887788609dc90b00aafab3"),
		expected:           decodeHexString(c, "7d6b367526b2fca36b1843a5f743cfef02e1c49f16071d1501f91ffbd24a03f5e0e64fbf550b6c366f5c0592bc83bdffa5db64897bfc38425354bea3e178261d"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_3(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "5c1c987625e91a8da14d9a4827e83e0a18b7154cfcd2a2e6"),
		nonce:              decodeHexString(c, "e75d8738acef17ff94a4809eded5cc74"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "4dc9ca6bf7c4e914ccb785639c80b08538e6859ae160b1c1"),
		expected:           decodeHexString(c, "d23c44c89cb4cf3a619df80df76c7525d37f83be93f723f55f26639969de8b12c687f17b4d4c947e6a53de7c0f7cdddef621502162a34cee70d51e14e79d853c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_4(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "cc15b70589b3e47b00a08e171860cd244355da6ad8e05cac"),
		nonce:              decodeHexString(c, "256503309468955f408de1a18819e98c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "00d394ec9e37c158b2b65f43f69112e6b75d6b1995d56a2c"),
		expected:           decodeHexString(c, "6cc2805953bc5cfb9a26808b10934001cfeb91a37a2f7bb8a91413500d6962020b1ae0d1ead603899eb234a51bebea221a803e1cec0f9cfd650686ee14b0c5c3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_5(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "f9b2fdd2edaa5f94efdcc3dc97bafda014210d62be329c18"),
		nonce:              decodeHexString(c, "99868805c0c36bd268650a48c6c84083"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "95e44007d9ba0beb99f8220852f86fb38ca63331f89915ec"),
		expected:           decodeHexString(c, "f6538f8f691adaa117938efc81885f56ffd9b421037b924be2b6e23d3851e4650e03f0dd7a44afa41a460ccedc3b6596dcb049e989564b0ff5ef6adcea82ab01"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_6(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "a97920820f9869a43875e3c2c6465ae5f8b0976ec8b6f656"),
		nonce:              decodeHexString(c, "927c2aad29091d96a45f28172530ed77"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "dbdb21298a359c9cdf820a04b286d87cc0a07496d9a49a2e"),
		expected:           decodeHexString(c, "de33efd68a83c91adcdfe9be0197a7912b0ff0c7ff4d2fa38e6b1fe427ce77775dd8a66e64060a8e82fd3afef9a73dea82a54a554a20ae3140c16b8ee3389534"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_7(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "94855425e828ce1f2e051114a90fa1ea86dafd8f2aa58eba"),
		nonce:              decodeHexString(c, "979a8cdbadf844665b13ad19769da909"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "695a5cd79866a72e2bd2b4b61be0708a4ea79580494d1bd5"),
		expected:           decodeHexString(c, "4b9f910556cfb45d3b3a0ba945e272d81e361597e3a26063b3862ec5ba6a5e9ee7c7d770e7bbce703e6eeb68d6eb12735a196cd201b5bc9f2ceb304ba64b665c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_8(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "f1a752eb2b869323484864e2285c08f23bc51ced5510c053"),
		nonce:              decodeHexString(c, "cbfb3885ad9eb3ddb2885329030c20b3"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "1abe3ae34c0a9b5df60fc7af7a8f65c3b765709c2fceed3b"),
		expected:           decodeHexString(c, "75a3e6e4517a3026f11fc6191a328abdf9166bf5e90010c1e8332e4a593e4e4931d6af129ee250d5654c6a9b601be6eed2d67119a7cd22221e68dbb9d713d138"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_9(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "eac15d4187db3c656f3ca82cadf3b822e792abd77d076209"),
		nonce:              decodeHexString(c, "0794cd4c856c3b06e5a08c01a2f7339b"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "df9f05099f0492b285d6791db70389a412af2e9bb49718d9"),
		expected:           decodeHexString(c, "3895e8b2751cac3dcbaa9c06b98d4bb48fc7dff175adb4fe9e8dde5e761f9c5f9d7c92c4e8320c9d6fd27b48713b39012b67154c15ea0be1d0b9b4523cd82e52"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_10(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "8ec770275486316465f6dd0a82d5a5f7529392ca59dae572"),
		nonce:              decodeHexString(c, "68e5b6bef9b22184872504b2b8b1bc63"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "066bbab885dbec60a4005bdf1b26ae2405b0f089cd72b470"),
		expected:           decodeHexString(c, "e237d14fa1de2db1eeed8d5612e64011612b929a59558eb11afae1e8ab010f9acaf6970b8612d76cb982bf8bd0597a6cf1450f5af46048f3768896540518d36e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_11(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "4d8f48add2e43f33e97bf4a047c0fbf4b69fc61695102625"),
		nonce:              decodeHexString(c, "7fb4e4b5e5c02582262f9c1ae6e3214e"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "387a8888565927a2262251d13718e897da8ddc41a595a1c8"),
		expected:           decodeHexString(c, "2b5c7c6d839e83f56e6e40e06ff17fc22eb1250dd3a1da76f2d6ed8b23696f2c1062dd1e64aea763e455835dbe68cb58ccce0ac26101a3edf0d6560a57be2f28"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_12(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "e2a30fa3cb3e33b97d83d0d3aa2df405377291e9920f3b26"),
		nonce:              decodeHexString(c, "41a5e74c8f196a5da2b4d9043e24cd42"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "15f08637daa6579270d87b0c7beb81c320981215fd58a106"),
		expected:           decodeHexString(c, "aa1c37503fea02a1c9d7724bc28b824246bf620329ba2f8b4d3aa99a8b5ad3aa4738d5b639a9e1734e52087f4da10f7c293154923b608190a29e6a0df76e562f"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_13(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "91f9697bb5086f3315f5ab85de1e1fb3c1cfa843dbae5a0e"),
		nonce:              decodeHexString(c, "a2e5ca730077769e3ab652b503830247"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "fe304716bfe5a4f0c612946f52b3bdea06f912d2231524d2"),
		expected:           decodeHexString(c, "fcd94f138b88aa2e498f7552345b253aa6ffe0bc159880379fe739c36092294de28d7fa7bbb574266f5005998b1ad1f2942b0e8ff18cb6b37e2f0cbf45c0301a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed32_AES192_14(c *C) {
	s.testCTRAfterReseed32_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "f3922cf422174ce26b0b691df1c8d44efd46aaefa620c204"),
		nonce:              decodeHexString(c, "54290f5202de671c30d7819755257494"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "c7dbf7f1161d3551c3b6e360ca732131c04fd5e1a43f9aff"),
		expected:           decodeHexString(c, "114eb92d66786679dcc9ea1afa04e4bafb83ab94f338b3bdc40ff66bebf1c46f736d6a912ac3e3e79c5b0e14927fbefb59cd805ad5515c501923dc47d6456e64"),
	})
}

func (s *drbgSuite) testCTRAfterReseed34_AES192(c *C, data *testData) {
	s.testCTRAfterReseed(c, 24, data)
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_0(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "35811473d105a6ae332bf72aa98a443ba97da55badb2e3c3"),
		nonce:              decodeHexString(c, "385bd442b5b7210c2d66109f88ac1ed1"),
		personalization:    decodeHexString(c, "16ed4ee59401550c67cdab99620257e54cfc7aa312c6814f81352123e4a28f11"),
		entropyInputReseed: decodeHexString(c, "a835e0140b52ae14df3343b65b110192baafd17ec2bcb10a"),
		expected:           decodeHexString(c, "4b8afb7c20bf941db7fb2cac02b46a45313334c04034b7e411b3607e19fc921dca47f19c5877e92086547cc6f1158ca4cfd62001f7e0f3af8a62e3c9888bf9ad"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_1(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "e3cff7142a7f220186032d9ae1710ada9214773a884f3ec2"),
		nonce:              decodeHexString(c, "1d827a2df68e9f64b6dedf0998e1123c"),
		personalization:    decodeHexString(c, "ea14febc4cc33846d266961d2d97fbcaba58c848bb828bef6336ab069a87daef"),
		entropyInputReseed: decodeHexString(c, "bdbdf2287bec536be16d62180d667ccfebd37d445920c696"),
		expected:           decodeHexString(c, "16a46ae5dae5ea51c58a8ff6eb43196de26ac97fb432c6333a418268f4befe2b7534a89beec1a77f239acd4c63367ad7adf9bc6cf2e4f742efccdef62624572b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_2(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "0cb4fa325fe328385680a2053a47194a7d6d201a028f147a"),
		nonce:              decodeHexString(c, "c3a712ec12b6ed9a214707a5f15a8b4a"),
		personalization:    decodeHexString(c, "5d4433d29bac8c03a1616df4319ff31f1ad62d76461af6d3f3be537dcd9b1a21"),
		entropyInputReseed: decodeHexString(c, "616ebb63cc14394809d5ce1148e85c26def9aaccf70b2f13"),
		expected:           decodeHexString(c, "2e0b663a5aafe0fc615ef4efa6be0c7ba1c474360df210ed0bc76ff458a3bd732d3dec7b890179bce2698c285383dc58045a9579fc724f9a24ca768a43b6f5dd"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_3(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "c9d89d88c933033a275309d3e8e484b92f3c22b1f2e5e084"),
		nonce:              decodeHexString(c, "492d6f25492d7e4334e72bfcfc2f4ec5"),
		personalization:    decodeHexString(c, "e1c99f116057842d6306db327b436a653ce76f4a30c308921366bfccfacedc98"),
		entropyInputReseed: decodeHexString(c, "8335fabd369e9dbbc4e2d634b37461e30a01db6ed482473b"),
		expected:           decodeHexString(c, "5f4aaf8b5a3a1a031448f949c4f82c501ad6b0b8291906ee005c1b771144c1c0daf5857a2564f7e0ac21df5d523443aba951ee5351a2e660e4d527dc3f320ca4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_4(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "89d5f49b38c2511dab84dd379d99965d897f82952947ade4"),
		nonce:              decodeHexString(c, "de6d0838f52995262fe8f91c2455a09d"),
		personalization:    decodeHexString(c, "9471587311f59db5d7b57b7b74db62687d85a08c2255f4a84e6c096e53b33818"),
		entropyInputReseed: decodeHexString(c, "fef4bf584b767cdfa297f07717e03b7e9a29f161d105040b"),
		expected:           decodeHexString(c, "b23d19fdad4fb8e45b2fb5ed807ba3acb11d56649ac2ef7147c8c78b134c5bba87a710cfb6e6460281dd4db7d6283e154afa164995af0249119838b22b8a5a4b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_5(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "3778c992778f16f48e46f68ce00be024696aaee2d53cc6c9"),
		nonce:              decodeHexString(c, "9ea6451ee31df30e60adbbf4d8298747"),
		personalization:    decodeHexString(c, "5aed8056fe05892e9c99c60f0bc1d9b24848f72a1722fc4f489a1ec20b5d19d2"),
		entropyInputReseed: decodeHexString(c, "421f1b05f47a0a1100ebb5ccb04dc3178799d2401f719859"),
		expected:           decodeHexString(c, "b584043cad7cd346dfd058d46b34f0ba0ca6c34b6198fff8c8614710483929a18baee5a67b1b648e6b27a7e9100cb28d1a942360f852295f33a8f8cf09b829cb"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_6(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "74ed969a5f8067a4bee99c248d799a8549479322460e1c6d"),
		nonce:              decodeHexString(c, "a8a5c874cddc73abd8a6cc5a0cb28718"),
		personalization:    decodeHexString(c, "d7b64d2dd3bdf5ba98a6f69e9045f261aae32ea4ace5ca054969eb3133270cbf"),
		entropyInputReseed: decodeHexString(c, "be8802ae995aea50a6de7dab9972b8653268affb687b46b3"),
		expected:           decodeHexString(c, "88754997c474bb30e9f147273f653c02693be2c2e47953b4c8ff2c37a1e1955219023d5aae09a8d39d772bdc583610128b89440c9d51566da8bde92cc26521cc"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_7(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "9ff4b4fe4bee71fcc83ccbaab7375fb20d51800e946aa20f"),
		nonce:              decodeHexString(c, "cc4aab82ec3bc4588b848c41293a5718"),
		personalization:    decodeHexString(c, "0f492617a189e9d535ca7b64adbb6a6dab3a5be513f86742f961d0c8daa53837"),
		entropyInputReseed: decodeHexString(c, "99d688499f6be5da36b9c9ba943eec847b4b6b3cf31840d0"),
		expected:           decodeHexString(c, "3c4c9a4c1c82f16a3f72c21ad9c0257b5c69fbd12b70a3a8bd2ca42a2033a231bec4d93246d25caa68d60c7aab4194381a309a1aaa079893f78a96accf1e10ba"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_8(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "eef4ce6412c029dbd0090f6bfb5642689c5366ef4b31089c"),
		nonce:              decodeHexString(c, "01f980773c1c7aefa218a767ae023ee7"),
		personalization:    decodeHexString(c, "a9ff7203ce16a720b8b6514dff587f54eb28323669d01b9e55fcd34a98dcad90"),
		entropyInputReseed: decodeHexString(c, "1926a729cc54e22869fad80db774d84f42d738d3c8848d1e"),
		expected:           decodeHexString(c, "5223f269daf6b478cd3105b8d689cb0b32247a4764d8b052be2c3fdd8a610ac893822820acba53c0368518ef3a3fdf98f5f4440789fc2bf4b1e999b384fb8a5b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_9(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "640cb52bb0720ecf4d69d9037c8a96999d06221aaa53f0d5"),
		nonce:              decodeHexString(c, "d8353dc33b33287b7ffd9100bccb5554"),
		personalization:    decodeHexString(c, "989f1ca7d63ad2bf8a2c34921ad7bf8af3d830dc1caea6a84d2c88bef3318aa8"),
		entropyInputReseed: decodeHexString(c, "865f772822f18563e9747154eca2b86f3866c3e152293e27"),
		expected:           decodeHexString(c, "eb3dfbcaaf41dba5b4cb1334df47b088d34e7e6e1340645d196b40579aaaed944e93e2de781b62152d75e82af9036d7ccaf3f83d67465cce2746f7b6b3e26e58"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_10(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "5e2676dc7c0e85e72aa4838639967e04f7701c111504a4ec"),
		nonce:              decodeHexString(c, "293a645be43ee792df89ea78c8fce575"),
		personalization:    decodeHexString(c, "ccb13adcbdf112979ccfcd1cb40f7443b08207da543836a6952149bc9a8b9537"),
		entropyInputReseed: decodeHexString(c, "640e6edb683fcba5b10d3c395b19bfe03b6498e140302a5a"),
		expected:           decodeHexString(c, "f70e60471a48049a4a85191f4dc2a5aa0da11a4551b3c1ee491367b85d4e10014bd914223ea70b2099d83165feef3e966e2972c554c397d33bfe5c7734ec3faf"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_11(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "8792c5a2c84b1b9dcd2a7fcc4a38d6cc852434de18b0f2d6"),
		nonce:              decodeHexString(c, "01132411a6df2f1fac374fba2e961d52"),
		personalization:    decodeHexString(c, "e985ad135f44c4f068c19ba42a474a782c4b2392a5a943e629802b2758231e74"),
		entropyInputReseed: decodeHexString(c, "51ea6900e3d2789e9f135354b6638ff3dd4957f92891ba66"),
		expected:           decodeHexString(c, "b384b72930d1b531b5e28221339ed28d045c7293673d7365ab2b0b28af0d82ca97b623f91b0b6508fa3fec78e2c5600740357e77c1ac673bd38612aaa91b46e3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_12(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "5572c0f3520ce25b4bacbd8276beaa9c8ad4f8cddafd8daa"),
		nonce:              decodeHexString(c, "46f24e7a1de769484464ab419721263e"),
		personalization:    decodeHexString(c, "dee6589c8b972b0b581852a26301cf5c2e7b582f5d574d4134db5a4620271045"),
		entropyInputReseed: decodeHexString(c, "f25ac7a1d9fba574f55a738d81d730d4a678c922abff5667"),
		expected:           decodeHexString(c, "bc7e16f8c6432da6d7f9f187650a330860e6b707c2fe2bda146a5adbbbacd305c81366e15f269fcf33b40cc1710e2efe3d96919b5009fbadc0266ddee85e02f5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_13(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "89bbe5d001db9fb2198db6d925d6ca12f68a2a5250a142a4"),
		nonce:              decodeHexString(c, "f458127b7a5dce83eeb87be824e1a5ac"),
		personalization:    decodeHexString(c, "2396a71468435c635e51f4f7d38c4a1157240ed9a4f6323e13e7db340aba1662"),
		entropyInputReseed: decodeHexString(c, "e13462325502906002ebf90bc2a42bb106f2d735f1d69710"),
		expected:           decodeHexString(c, "0cf5a25485ac140bc64971d9b1734fa57badada8ea2baac70bb29dc23ab53279ebe5444de23198796398214483e8ba12ee647c3f0668dfff2d979ae04143fb3d"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed34_AES192_14(c *C) {
	s.testCTRAfterReseed34_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "8df6796aa35595be1f4f9f0740e2a822332c82fbb9ce492e"),
		nonce:              decodeHexString(c, "125505cb0c37e04bb9ac58990699a925"),
		personalization:    decodeHexString(c, "ffd75a98dcf5218a81f5820087eec84a4a076001b01e24cc8c156803ccf99435"),
		entropyInputReseed: decodeHexString(c, "599488019a36132002367df6e2a6f82735ed55d5fa1c9146"),
		expected:           decodeHexString(c, "73b09aef6783b0c9d308dab8c961be7ae34affd4d0725684298a2c7657e5a33fe367781ee1a7f27e337c23ef7a355d7800f08542005f9686f2a8762ee8f6db35"),
	})
}

func (s *drbgSuite) testCTRAfterReseed36_AES192(c *C, data *testData) {
	s.testCTRAfterReseed(c, 24, data)
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_0(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "ff25661b6b585c5a31217a9b5c20a6e307f70b126fda2e25"),
		nonce:              decodeHexString(c, "1be8f51afebd48145541603df92e5d0d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "29855dfe13480058562d337e16ae0c8753cc4eb5420c8825"),
		expected:           decodeHexString(c, "8607c9d784548f2f37f2616b244e9f27a30092df9424c47b3464862e675f03d4ec6cd5ff79f9f4a5d388a603e7495d39477955460ac2ee0c2ce4d3d834ef5174"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_1(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "39af29f392bc317898a66a84ad843b9e176987b029d0012d"),
		nonce:              decodeHexString(c, "efa622ecfecc7b5293768b48dbf3e5a7"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "b666306e4f02f4cd455c3a799438733677c5965f482e325a"),
		expected:           decodeHexString(c, "b736c6ba430b051fa2096ee7f077b0ceda0b40d4f397b4dd559c294a73003df380b266d2bba4eb94acc62f45987126f778c94c0208a330edc8bfa3aec7dd866d"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_2(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "3430be27583710618822a1d8b4a0fdfeabc12f8d2f702387"),
		nonce:              decodeHexString(c, "c2bdf87547fe9cd7cad23bc1f63049bd"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "53e2ec40c031d860b77a00c86d1b6590a4bd02f7b27b1ef1"),
		expected:           decodeHexString(c, "edf1b1b648907abe36e5058fff3ff7aa6682e01d05815617472cc50fcf17c3ebe0c79c0300499e881d106739e6e4d6f40221a77a242dcde29687a2da7e181209"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_3(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "27f1d6e2cf04bd6e97f1ac028a7f07687fe8a017447802dc"),
		nonce:              decodeHexString(c, "b3c49d6352848c4f72d7f82e5c3572dd"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "dfb13a1426bd48defcfa7a91e13467eeb243075c23a6a419"),
		expected:           decodeHexString(c, "7d904e4ce29eda570c4ca27af38235af666a251a4518f064963d8c3641806e5e7d963d1f32090b8ae8699182b9ecd625ec61e09c47083488eab75656c2bee4cf"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_4(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "df28a9331b889913657c59f9546cb6e33393ea65858432b1"),
		nonce:              decodeHexString(c, "d6d8583bb7ef495a5b108730af2cf14b"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "7b144e638de53a661817c295619ea0e08f2bcf3fe1c78290"),
		expected:           decodeHexString(c, "ea760ab5da811a6b1a179ccf531902e15fba08014ff94d199a1cd8b7527487e43c36f6874ce35211e6d719fabf8c0a05d4f3782e78f9f8e89131fe9f3921c8ea"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_5(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "3ca78abb9347fdf876d298a37de1baf0679f9e9ed691dd63"),
		nonce:              decodeHexString(c, "c5b357e7f277e44b2aa52c6be5ef12ed"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "968f83030b805037d2bd5db9eaeebb5c5bb4650b21afec1f"),
		expected:           decodeHexString(c, "d2a032f77b57d744067667cff18840d034e27381943ef141fb947f34076750b9e90c2c033eaa9de78e48fdb1926e3124c21078fccc44e35098b489477522671e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_6(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "891cd44b246f6cc47901927a0174e225078993b5b64d8972"),
		nonce:              decodeHexString(c, "fb4f2eb73c303aa6e3d612a91f251930"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "5ef1972373ebdb4c7b414c69845a8eada0734ea46fa1d012"),
		expected:           decodeHexString(c, "59e6f2e416e0fa43b65fec2bec3d511d97e0b0d537ebf84057baca2950449be1253f584d575403fb172033018efcc094e43db99840881d944b450afc6b9e99c3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_7(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "55709ef0148e66a605e43e76ccd1aa9ce167e25385716cc4"),
		nonce:              decodeHexString(c, "6cc4711ac143aac7363b9f5fe116d550"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "b178ff732e79459db512863f65487d08e0185c3e825f7496"),
		expected:           decodeHexString(c, "0d27a39e961057fc0303818c800587254cc4a7d7da33f216a1b2804a2785d722aedef0e16c02973007383cb7574c64ed1f8c528cc8b0bddeef7707cdab7b7ea6"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_8(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "f97b442ef489f822004697ceadacfc240e533db43cc9d21a"),
		nonce:              decodeHexString(c, "eecf1b9a52c370f0fcd64b5d2d7b053c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "e97b2617e87c6ab93ffb9ff06408e444452e06d01dbd9949"),
		expected:           decodeHexString(c, "2f82dff54324195ae39e92630af541b39ebbeea0adab967fe51597c8496cd6637ef74836b7a92d9c7f699346a8b4a8907870409fa5a2a39bf171b8d7e5806c9a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_9(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "c5d2d836c375633819231bf91f50f8b7498c7dad848473e1"),
		nonce:              decodeHexString(c, "4525025b4e848f81f4d4c67b5fe22b67"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "e9f7d8d354cba1691c40ebce00a2dd73c4569d793e838a23"),
		expected:           decodeHexString(c, "3ae7a6d76787abd470d2a9586549fe0815da690402c83d4d7457a02c1d795c61c942df128ca7c2fd2c0589f9a0c9c08b247579ddf265af06affa9f0441f69771"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_10(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "93dfc025a3a839a77d63fde80c3dd28d82d0d2999d19d04c"),
		nonce:              decodeHexString(c, "97bebab36846536f60d315fe884347ec"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "b415c70c93ed68964c9f152270423cb382007a5ba88827a2"),
		expected:           decodeHexString(c, "5aa70288c848e1c2ba058d545e97fd9102fad25fb609da1e301c252e4ff0a7eb537f3c411fd048566764bab617060465bc6c3a2ce0670e68926ccbf4e98d9140"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_11(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "9463fc03c17bdb566f1408cf6d052669816fd79d4c550286"),
		nonce:              decodeHexString(c, "3142c030d33f4651548c569b64fb2d45"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "6e6d9a7167fdab95921d78fe197d065ce7bd168975473c54"),
		expected:           decodeHexString(c, "d182bf85d948b2c8b05cd2a98b2c7baf4618c12387b41d72babbc9f92b9fe3e4ba6552252574c40459bfe74d9e790b294911ca0777d2d2c4d54929700699ea0c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_12(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "ccc5cef1e1774807b2088ea248691361b5603806ce613109"),
		nonce:              decodeHexString(c, "0256af445465dd03756c3001447eea4f"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "b6bba417ab2c32ae8121c9b0af7724eeab15b7064d9efa91"),
		expected:           decodeHexString(c, "dbdf755f7a3f7e8ae10da9957fd6f068b24d2b253818b8fd1170fe17c6eb98112d56394559f50753761b4860493264646c517abcf5c86a29cd1acd6d1451e394"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_13(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "95b7c27807116cbb6ad8b69c8895ef115bf56587d8c00f6a"),
		nonce:              decodeHexString(c, "c947784114f6f5ba965286fd9b31f3c7"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "f336b1ab1a4febac39ebceb2ed919a81b4388b13ae777242"),
		expected:           decodeHexString(c, "bbd70942bb2772a1c0097bc5cce01bfb64dc27cc20037f859e23a6c78575aae4e2eda38c64599a40e706ee0c119fc1eb0e139ed1a02f0da9d7cf7b84cc7e6934"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed36_AES192_14(c *C) {
	s.testCTRAfterReseed36_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "770b80e15a3eee89f75608a780934b159de6ad3f5d78b6df"),
		nonce:              decodeHexString(c, "d74192d50f0f739094808987ed94d382"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "dd3e0ae94de73952665cd75c729404329098d610940a9a40"),
		expected:           decodeHexString(c, "111c1575c8eec37a4272f59e839b6d6da24886a82f1d3389952b5f5fe9ea6b352e8473959dfd617d0cd8588010a1f5da8c9a9dd036610a93d4e940e09a95b164"),
	})
}

func (s *drbgSuite) testCTRAfterReseed38_AES192(c *C, data *testData) {
	s.testCTRAfterReseed(c, 24, data)
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_0(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "a7118c155d69705519f1a56f3f524efcb32aeb32a98ae81a"),
		nonce:              decodeHexString(c, "0207f3534f118c71cc11f181f5c6fc0f"),
		personalization:    decodeHexString(c, "8d11491c83dede012f271c263f8d4246c3efdfdd288710c67e979f4fbbc864cd"),
		entropyInputReseed: decodeHexString(c, "66abef82f794f35ed94282deef0dfc228ca5ec83257e81d7"),
		expected:           decodeHexString(c, "5ddd46baf5c6a5e30943c5bda12f5db8ec19c537f1702afea31292b4a9a8d425763a9f92b36f616f4ffdb9160774d8776433b7c05c46fe6f66c403736a044be5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_1(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "45c2af2d3615d4493beb79220c2f782ba5642ef44547147d"),
		nonce:              decodeHexString(c, "af672d645a6834d6565fb32d15d4085c"),
		personalization:    decodeHexString(c, "0758c24fc92f126b41ea906da2bd9e8f9732703b13c4f2334c96fb55bdbea998"),
		entropyInputReseed: decodeHexString(c, "8dd15d944cbbcc5413996d321620fa1264789534bd228fc1"),
		expected:           decodeHexString(c, "6e4c6cf2192c959376e1d9f3fef0a715a9932b1cad08b9c653e7b1acee5313051fe2f43cdbe0f16a565039d547cab6de9bbcbe3af12d9057a9bc3feb98d7210e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_2(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "5fe4905c810fccffd3e2985f2094690d7a581f3e5ad39971"),
		nonce:              decodeHexString(c, "3a12f5f2d90627b9b7212c33bff3de9a"),
		personalization:    decodeHexString(c, "b4591dc7ef2705808edde80a37f625705e5ba2ab0d4c794f10ac0df6353f5c15"),
		entropyInputReseed: decodeHexString(c, "79d3f08470e18909e85a0410c7e4855b155cd59b2cecc47a"),
		expected:           decodeHexString(c, "c497be7eae974f2ed65dfb0f8b6cbb3ce05db82ff31f36f04aaa44e19b0828d0c5fdb59cf588529e5f7eb048c42d503e5be7372657347e9d3edbded34de31633"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_3(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "7223175b57e3e3a22b95bede4b62b246eecb217ef87147a4"),
		nonce:              decodeHexString(c, "ed25c6f4c4879fa99f175b2e470c8dee"),
		personalization:    decodeHexString(c, "ad21bf9e96a91b56fe9b199a16a674885abf223483c76c6c279ecb253dc5800b"),
		entropyInputReseed: decodeHexString(c, "544bddff07fb6246fecf5aa4ec7a88350afe2758240bc558"),
		expected:           decodeHexString(c, "3e52224e4891a17c1019d28d0c90b0365928660f8af1f66fabd6e687d81ca724f134014c7259d53fe795abc34b6fdc039d9cf2d246cc738105e17bd554752f73"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_4(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "1deb3128bd87022dd2ff7abeaf7bfdac66140c319f1dd24d"),
		nonce:              decodeHexString(c, "d6504f4dbbe2feccd89c9c225cef0c7c"),
		personalization:    decodeHexString(c, "266737c2c2825f2d0c8539757ff8f9e85ad94fa4ec8e2a2958025f0fc2c66487"),
		entropyInputReseed: decodeHexString(c, "f9679a220b82af90385c5e04fcf891caca1246c0ece083fc"),
		expected:           decodeHexString(c, "9d4194c988a7bf40d47d406094b7275340a6474357c8f4ee5fcedb29f5c3c97bb9f6dc957a4ce3ce541fe43c4d2c547c85e4823288af9d47ba8d4f89d3d46caf"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_5(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "92295be38766b648cf442b84cd3eaf02282e378c00d3e42f"),
		nonce:              decodeHexString(c, "8f1b227ab2cc6564868d7a219b36f5cc"),
		personalization:    decodeHexString(c, "b572b2bef84403d8a288f0a5475c00f186941f7288023b8cdaca3ee65c4b5e46"),
		entropyInputReseed: decodeHexString(c, "a5b8b6f44d8aa23b6621888545098e2554a68115db1b3f0d"),
		expected:           decodeHexString(c, "a11e67a353a2e64ee0608eea12e6db0b35c1c6d9ebf8869bc4ea4806009b423c68c2f6b39130a082ed6b063ad57bc2d75569b221647c9bd1c038a73604df2cd5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_6(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "0b6ea66e9619c8d3e1bd52b365628a4a0b7f948372df2853"),
		nonce:              decodeHexString(c, "6cbaf02c173d739b0094c0a25ef20a54"),
		personalization:    decodeHexString(c, "2d23dc5840f8fdcc8a8b93ce301418faa817ecc6612ff0c91b01ad7e22a864e9"),
		entropyInputReseed: decodeHexString(c, "edf76385d2c99516988e91869e09e1c07e655655428474d4"),
		expected:           decodeHexString(c, "1757ead14a22c4ce8cd190fda36710c3ee8056268bd6ef448bc38a34057daa0d55e709bb20fc72bbb5d4c1fa16b404269f9468c3029ad44448005d5be0f45526"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_7(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "d010003f835ece36be84e253ebb829b6689ae05582da399a"),
		nonce:              decodeHexString(c, "17281f73427d331f159dc5cca3b84269"),
		personalization:    decodeHexString(c, "9a9c9540e6f177fde40ada55cf6dbcb53cf30052bf69fbfd104195db59699574"),
		entropyInputReseed: decodeHexString(c, "6b72933610d8efdab5b23656a657f6934592f5c86e018b4d"),
		expected:           decodeHexString(c, "aef5ae882ccb2faecff84c6484248fc13ed80791e8926675a9c01184d843e0aa154d97dd1759e4efb1788b0ec8fbcdd6e59e9f1db02ad47c8b16781fe813c250"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_8(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "a9f408a0a97285f0c53416e60df55ad46a1bee91cc9d4636"),
		nonce:              decodeHexString(c, "73c7853fa67681e4c25587b547110eb2"),
		personalization:    decodeHexString(c, "3ae26b8b2ebd1a15316100969d335a383945fdf6bf9f7c1c658311811679933f"),
		entropyInputReseed: decodeHexString(c, "877bf6c91d6345f777dc1df5650c1fe40f74cade7b3b29a3"),
		expected:           decodeHexString(c, "679fdcc2c5b10bcfee0ab322a1a3fb43a830039e9c23574fc6ee772cfe6ef7991d350e52c0cbd7b6b3e4edb2abd3bffff6d722bbae43305631df0abd109d6bdd"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_9(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "db355d726c79eac6d62258ed92d4f31cd3e6a8e0cda05274"),
		nonce:              decodeHexString(c, "af52b3f0d43e3a0310da57b4d0c6a284"),
		personalization:    decodeHexString(c, "51d6b8403b7a3729c5e05d595b6f3850d220b233f4ec3052ab678f4a3e1ab429"),
		entropyInputReseed: decodeHexString(c, "9f1e5762f5c75b1ac98cb068817739feed86facd98fe85b1"),
		expected:           decodeHexString(c, "5e6af300a4ff0c0042d1e8c51b0eac14902b9594db95c35dac8eb53149663c6ad59f9f475c6802b11aa094494413ffc61aca806a4557208b0c863b0c10a60cdf"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_10(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "4dc65abe7edc6628a9ce20fecf1ca9cc11e106c247aadc78"),
		nonce:              decodeHexString(c, "269e93bad38ff69af72330e5fb2960b1"),
		personalization:    decodeHexString(c, "5e10578bea14b9806c4720f199de43d4b3498745d367ce6aea805d176b83af85"),
		entropyInputReseed: decodeHexString(c, "ecf4f2a515794d3f4e198246f772645d341214a8e273ddab"),
		expected:           decodeHexString(c, "41335b222f063d9f1e8c09b5b1fdf9d20c8d097593917a55327f8128f5bd8abd42a2a89a86d10ad636ea8ddc0aadf479c2e6e6786ee166d337d8cb3263d6b3fc"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_11(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "b986276f9fbb11abc12a1471beb82cf5e8ccc53c02dc3527"),
		nonce:              decodeHexString(c, "165b23a4478487e0a4efb40cd3f86353"),
		personalization:    decodeHexString(c, "a820fc958868894bbaaadb333aec48d2f9ca8eb33780b217ed3bb8aab792ddc3"),
		entropyInputReseed: decodeHexString(c, "10c4476237b23ce6f5ff6bf2690cf86a9e0a5deaa2ec945a"),
		expected:           decodeHexString(c, "ce03a47af0f4deaa7ba0f2ac332d4c467dd1bea4de2791abca975543ae78560967e675e90d1ecdb6643f39aa753b8619b9e2a76e4639e62702dea23c6e609b71"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_12(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "d6465c9ab55c68193f5be87f102408e22bdc9b2981566ff9"),
		nonce:              decodeHexString(c, "0c0bf9bd555b915642693ed45692b2e3"),
		personalization:    decodeHexString(c, "5898c4999a56758f679194a91d88aa1ffea861743cbf99fbf53ff8eb91aca185"),
		entropyInputReseed: decodeHexString(c, "8283c48d8f7750ad74631197a58b30f2a536bed26063ff2e"),
		expected:           decodeHexString(c, "2f1f484fa276476531c09e2b8a47454ff4a74b0423793efa056b415bc853590076d3277369a728aee66c1f3bad913f33b482cb8e0e805891e20acf14644dd55a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_13(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "1eb696f5c64a3c492185be787989a3e014e8d55257f39e1e"),
		nonce:              decodeHexString(c, "0d7b8a3d6bee621d2a29742c058e8666"),
		personalization:    decodeHexString(c, "9c1484ee471c8847da3d97acef0d0283af1d95f735717fc231c12846bd8eae27"),
		entropyInputReseed: decodeHexString(c, "17e333bbbf9faa7411cf13dbb540e05be954b81715b09c68"),
		expected:           decodeHexString(c, "544c5096e69f2b3bf74d8ea789a9572d2e61aa836b2c5aa5b02bfe9edd127ca1cc49917d8e7fb815497096857cce595c84419389662f1bc842b5fbb46d44dcd4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed38_AES192_14(c *C) {
	s.testCTRAfterReseed38_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "0e14a05ffdc1a80a9ca9c7f377ced69c0642ab8490c692aa"),
		nonce:              decodeHexString(c, "25098e574502ed6349edc0c8d3a28006"),
		personalization:    decodeHexString(c, "1d0e83b9f32571af1bd94cb15e4eddf18a4e565204746d7a34d648ef17d386a8"),
		entropyInputReseed: decodeHexString(c, "3b2cdfc8e4eab283d611831ed121a9d1dd2c3b90c86e4298"),
		expected:           decodeHexString(c, "4a40d4b4ca4290c67b5222fed57ddbba72b6aa20a57048a32478b9941bda2e96e4f35b3285d512ea1ac0282ef2d4ce4c223a6dc79b911e1bd3576c9c9008f957"),
	})
}

func (s *drbgSuite) testCTRAfterReseed40_AES192(c *C, data *testData) {
	s.testCTRAfterReseed(c, 24, data)
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_0(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "8516a8788f6a58ddadedf8175aa289a2bc310c07ef0d0d82"),
		nonce:              decodeHexString(c, "5db7be3e068300b25bd94974475e58f5"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "423b9912e0e6fbef38642c46df7dbaee2197173352f15267"),
		expected:           decodeHexString(c, "c612ffdb9ea6e20901e8f58c180dc4d79b9a25bb51ffb75d7b4076f8f6791c232eeaabb2e43f309155e53810ae795eba667124a9f697f6bb356785edd9ff3959"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_1(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "8893593a21e2df41ddbd7fcee30018a69ffb2335860d18da"),
		nonce:              decodeHexString(c, "fd1dc585cc195c10d7e27418ff2ddd35"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "3c236be2707e388022e7307e1ef4de387b3a6eb95086b2a7"),
		expected:           decodeHexString(c, "3083054ea1e27687d8696b746298b3add9f3b56d7ac83fba75c472c3d5bad1bfd5c537e4211649f72beb0dea79f097c9eb13a7cef084333d9586663e0728b290"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_2(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "0ae9a60dc1f074b2dd86ba3fa2e329b9cadf7afcd94e64d1"),
		nonce:              decodeHexString(c, "28f3bb2c48ebe5829d31283024f70a55"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "d0ffbaf462afe1fc9b5c459328036e115d4fde0d0b87e287"),
		expected:           decodeHexString(c, "678285fd0cba3c39272f64adafc69707cbbbcc502643063b0cdafa32c45e7aca70ddfaedc656d743cbc980a185904697c4e4dd1a1026933e42598d44f3c1c99e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_3(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "7d0efd400fbc7ae6afb8b085291f22d9d6f41290c087bd24"),
		nonce:              decodeHexString(c, "11a6a453124adee33f8ebc6885ccafdd"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "33d555aebdc8c52933b58bbaa27c2b46bbc6c59ea9c326e0"),
		expected:           decodeHexString(c, "e405984862b40e21f29678c410daff864ed9e70ab41fa30782a73390fdddc3e09e0a896f12b1c85e6a6688203ba8bbdbacceb50fdf90a0d2d58e881ac94b3da6"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_4(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "7fbebdc48463e3ff813588a108c288e06ccebbbf3e0ab757"),
		nonce:              decodeHexString(c, "8492b8cf8ab3620ade2728c306e07c2d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "226d45b7d77cf74e9de74cb24749b926675dd8f9b3869455"),
		expected:           decodeHexString(c, "c56f7040ef814f244991bfb431ab18e1ceb6c7ab20cddb6e9cdb1d7740d6a14343d5e8dcca2c70f5cae2576ba575795bcba73679099cc03cc501dba6c97583a9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_5(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "e2e695343bcb5d7cbdedc3763d951551b54861558a8f3211"),
		nonce:              decodeHexString(c, "418739cd69cae50ec9c10357bf6ce08c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "d9362fce7bc48be437c9e9b670db331d682418f61d772500"),
		expected:           decodeHexString(c, "013942701901bd4dde2e383a5f998f66f20f9ad575bcd572903a9d8ee30e22acfbe0f263bd0db02d38a99473fd7dc89544cec1c4dbca81bf2d739b2148f1afd7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_6(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "1cd89cfc5597e809edf9e68a25f4411c9df3c3f57d7d7218"),
		nonce:              decodeHexString(c, "86dc6a4333fb03a89993cf5d369578d6"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "b65e35a579bfc37081938d696327a92f02a28a1b60865521"),
		expected:           decodeHexString(c, "2105537940b64a57021d435a1acf035cb307d24181c09d6fc67a73d6e36fba8a62870e726443524478a18d790cdebaf123f5ea882d3470c8ae3a1eb1dfdb4734"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_7(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "1cc3e8f1c33760c4a0fdda9da7dabaecab10d570c6dd6400"),
		nonce:              decodeHexString(c, "1b73adb66783bee024272b695b0e19c3"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "3bc148c43b54cb126a16e75da54f53897ec0a28ae589ac64"),
		expected:           decodeHexString(c, "5ed9d651921f35118a6fd8ac3484f9deee2d025e0ba4800efcf48e5a8e6322fe4b49361801c77e0268e1c5e55be53c8fb5668c7a289670cb26b5a85bc53361f9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_8(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "7219739f3e1f8c239cf07df899c87c79cb4d59e7b9ec8c64"),
		nonce:              decodeHexString(c, "4fb943b62bb2374ebee50b6967a3ac46"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "7557d7ddef0a0fa1d7b8ee1219f0a1be84dd0cab71d85b90"),
		expected:           decodeHexString(c, "8b4ae1df4e76bc754dee8e64a0ae0d0cd67a628e604f8e285bcf0e081b6b3dc11c4f43e85aaaea193562e135cf061a71fbfc469b7f2c03833748de3285bf9c2c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_9(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "6268c243d1c9094f63b1c45f4bb1c44f43ca28495494fd46"),
		nonce:              decodeHexString(c, "9fcdcb372d375441925d7d275b582d73"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "ecf4a3162aba1295a335d5e002109c95ce2010dcbbd1f8a8"),
		expected:           decodeHexString(c, "08a2aa608f79b3f5ee1713accde39149efb15041021fadc5ad2be4ca3b2ad824b32767c2eb35857bbf14ce9fdbdc301bd6ab51ec16e308d35ba94195fa51ee58"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_10(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "c905f960654bc6317cc5bd6b32ae7961b0fe5eebb4e9e5cf"),
		nonce:              decodeHexString(c, "c72baa07ffa45eb5666bbb01109a4f5f"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "ee91c518f4da928b36a5289d350d419702857e243e62a9b4"),
		expected:           decodeHexString(c, "0d742f74dc8f09c61c363a187948811e9a504afcb13d5e6f18e9f1905a05cc754008ac57e2fcdd4c3bdc536850831e7d1b2d499d95729ed57373caf0b41a559c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_11(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "7f0af19f0947d3bcc2270ad8864232750e3a79073e54aa5e"),
		nonce:              decodeHexString(c, "d0fe40374598d5ad63468aa06f92c503"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "658e55215a906334e6b8026f6fc36575995001c9dbb8f890"),
		expected:           decodeHexString(c, "a8ddbbef58ddeb3bdb923b64ae8f43f0aa3ffec303c04d0aca8ecf0d4525fed842155dfe4638bb0892087f8f7d28b128296d62e98a51f64932462acb6b841a82"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_12(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "0b8a7f72c3ddc63767d3368059853b2cc3d8309534d44c31"),
		nonce:              decodeHexString(c, "c5279d2522cb770a8207ec2d384f7ccc"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "c1104427c197bfaf5d4a1769fda012eaa2f54e94177a5ad5"),
		expected:           decodeHexString(c, "567d3c406e827e5c0f438a30417659d227f3ad1020a5ccfdaa3e64cc9fcf6b0eedf83b9f1fbc525140bc60434c0c02a3ce5482f1aff85f30c6f98dea4a22edf4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_13(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "0097b544374041cfb45727bfca769c7586cae52c8b945c18"),
		nonce:              decodeHexString(c, "0e9059b8655056f9c824180fe107ab55"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "03ddb3e6f27342538d210b02ce2bbd420a52d3142a69b0ca"),
		expected:           decodeHexString(c, "7aab29b7c40c29c4fc71566b3f9f6ebc37ea8cd9266d0c8b4d50d81b07959c228e0d96761aaa916568ecfebe6c7af3e9c02c5564004f225da91811c5b761419e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed40_AES192_14(c *C) {
	s.testCTRAfterReseed40_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "02e103c5dcb55e065e50f7f2660283b0e30872433cd36620"),
		nonce:              decodeHexString(c, "b2f21f7a101195d4a2e7dd7353773ae2"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "95e4c2f976931cb5036fae1d4d58384744c4ba3006fab366"),
		expected:           decodeHexString(c, "f04955bf052cd18899a448a25079bc03e3619c44448fc55b6e6f04c9872480378130d8d2ae353ef255b4a9f536c962e2da01fa44b7d9d34c3faf14b7257a62eb"),
	})
}

func (s *drbgSuite) testCTRAfterReseed42_AES192(c *C, data *testData) {
	s.testCTRAfterReseed(c, 24, data)
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_0(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "b4b05ce9ce074183a306bf07eed684bf065ebffad9758451"),
		nonce:              decodeHexString(c, "733a68b1d276f154e79b3dd083d2278f"),
		personalization:    decodeHexString(c, "9e884facfb9e9d4003471d762be01247fbcb18d9084973a2b833764d715c0cf8"),
		entropyInputReseed: decodeHexString(c, "be9003b06e75265c4fa54c31f6f5bfd21325e124fd541769"),
		expected:           decodeHexString(c, "38bf4694717062ef51121bcdf1dc700ecc1a018a2da407417a8f6311f197e7da9c7b4a300546a68c4a393576d1512e277375a0e05c537577867c937062d70fd7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_1(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "a4f6fc1a3dfecb1377699826b8f4b5129e584bf8d26b3a82"),
		nonce:              decodeHexString(c, "47f7a55a783870c5cb621d67387d6fdd"),
		personalization:    decodeHexString(c, "b016db1c111f203d108496dfdc799ee2ece6d7f56c8b73c9f5eb24e188a4882f"),
		entropyInputReseed: decodeHexString(c, "35b7e715ceb26be75c230d442ccead39c9ab941038d474cc"),
		expected:           decodeHexString(c, "e475262f8f34b0d23eaa753b0cdfe5f49bb54216f89c7b73cbd5b4cd8c882cf9f25b4059e39872c3a79020fe44bef20a39e3df165e0b6e169ad46d89ee3afa0b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_2(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "2c84f1a79a2cf4fac8740ed980c65732c5b257bb1c2cea36"),
		nonce:              decodeHexString(c, "40f1fdf6fa5ff03766ead7b71bcb44d2"),
		personalization:    decodeHexString(c, "4f8c7b89317a92fae18b6d68f7facd601c4d4dd539c3a9c3ebe289186ffb2f22"),
		entropyInputReseed: decodeHexString(c, "a65ba690cf9ba574d010552b7b3dfd388b5652f3233dbcd3"),
		expected:           decodeHexString(c, "ea068fb62d2f78b1b4b1925181a9f613964a17c3e1f8a9ced2a96fa377123c01b7f784774e940831c1f52dc15aca3bad24ace3422edf56dfc8952db861388322"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_3(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "263cb0780082abf9f3b37ff2b8f810db35f970f4749e1e4b"),
		nonce:              decodeHexString(c, "436558425d3a4326aada7eef7b13c1ba"),
		personalization:    decodeHexString(c, "30c82be25ebacba224aad9c8b7d4993e7bae908205ebd11c19f1b35b09d06af8"),
		entropyInputReseed: decodeHexString(c, "63d560f8a0eeab76ce62a16d125d392a36ff342b034ce8a5"),
		expected:           decodeHexString(c, "bbb39eaa24a98417b6afdcd56d1541e30545a9444cb7bc6649e9bb8b4b52bf3a10a7a20b5bda1e7640fa15ce0905d76ec492dbf3b3746589694acc2ede36b296"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_4(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "54255d158832501721af074f1a053f04c717c938c3226d7e"),
		nonce:              decodeHexString(c, "c8b9451433a29e92099982fe6ba9955a"),
		personalization:    decodeHexString(c, "c26c8dadf21051bf82bbd09d5a4550da1e90fc8d3b3bfc0ec2ca55efc5ce66ae"),
		entropyInputReseed: decodeHexString(c, "30035fa0ddf9bdbf2e09c69260ec318ca3e44d6080a28b7e"),
		expected:           decodeHexString(c, "84d15784b02087e26226c663c5c86d5f0617a1503010920cd029f7ca91bb82b6633050139eb423b86f2924d50da7787e1e47da02b8348591b7c7a7c61dbbb8f5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_5(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "253ffdfad4165e6e3f00b9735350a9e17ef2fa1a048308e6"),
		nonce:              decodeHexString(c, "77b959982f91a2f521e7e8c25550675d"),
		personalization:    decodeHexString(c, "16435f2ab0ef6ee1a16f9ba15ab5cf2641ab076b37e6421aef73380b4b896991"),
		entropyInputReseed: decodeHexString(c, "c069cf7c199146baa67e5d8023b562bde6185d3dab781211"),
		expected:           decodeHexString(c, "a58d58c99043c0e257fee38efd78281296d468bca2a91918d817d31b00bb52930879141de24bde701d8376e8ab07e843cd6c4e334e5bd0c36f92e38a19581992"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_6(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "cfd16ee953aa4989a208b9fa8e0044ce2e4ed6d672d182fe"),
		nonce:              decodeHexString(c, "de8e91e65814569ff04767aa75db6bf5"),
		personalization:    decodeHexString(c, "de223d6643e233190ee0218fef6305c142d2feaaa0abd6ac7c8bd895a00e9172"),
		entropyInputReseed: decodeHexString(c, "ad5a4fb9e03f8a51ab238235ed272798be43db4217a5b363"),
		expected:           decodeHexString(c, "1d59821954c5c4ad2248e0593eab775cb208a1887becf9cf93ea2b0258af6e3705b1d7227a1401cbcd3272edc4cf985a05c345492d1bb9e9047bbeec7aba2436"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_7(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "36d2236948400dea19142dd1c53c3e7db7c3bc7ef3b9e275"),
		nonce:              decodeHexString(c, "90a5b91030cb2537d4bcbbcec6cd4386"),
		personalization:    decodeHexString(c, "7cc51d67304005b5231c3a3ad3bf144eb6a174bf1a0d29ed97b887cf7c00c3f6"),
		entropyInputReseed: decodeHexString(c, "8307c028001269452b72b4134d853416365766b8b8c27eea"),
		expected:           decodeHexString(c, "b64ccba08e5d9700a6c7be681485f67cf1b2eaa7a79fcc0cab39d3b268e366147c9542d4446fa072614fdeb50e748c2fdd85f4d516863aa8820f69c070761c72"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_8(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "7eb189f1bcb547d837b7553231325b4ddae1ea87b938b0c4"),
		nonce:              decodeHexString(c, "9d2dc96e33373f89e4da9f5ba17eb1b7"),
		personalization:    decodeHexString(c, "1c56b2a0253409cca57383c95712bbabd5a55b329987d1bee5ff8237e7bd9e70"),
		entropyInputReseed: decodeHexString(c, "cd9b86c1b8b82c72f7cb2bb940ee3189a2554187dfa8fdd4"),
		expected:           decodeHexString(c, "e95dbf39d1482bd03a57735776ddb318e0f9ffb0ccb513607de3863b4fd0fb405e74d28b2ded6ee1972e2ee89a09b6641761fa18a2f6c123aee9f7191efe528c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_9(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "e061799eb8a570f6de95ba5f108fcc1d2c9ea16ccfc2bece"),
		nonce:              decodeHexString(c, "87710326b978b2f4343e2670b451ed31"),
		personalization:    decodeHexString(c, "27a3a92910368da0984aaf3c7e315bed2b6b193415d438a37def8f2ce542a323"),
		entropyInputReseed: decodeHexString(c, "3f3c153dcab7a008566c25d3fa71bc7104767c1d4ae5ca02"),
		expected:           decodeHexString(c, "6b75481f70da166f3055391c705ff6ec675dc912253a5de9f661f68adc097c64334616839c67f298aec4104b513f4970e61dcf3040877331c24d87eb091b9d91"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_10(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "ba8ca361ab65b9f156e8dc5d4e50c8ac4c7876cbeaae09ae"),
		nonce:              decodeHexString(c, "c216c9af4dd9440ae5191541b40b1d6b"),
		personalization:    decodeHexString(c, "25adc273f5d83be27023352b7a78bf0cbdaefa76725a45aa9bee06137351d24d"),
		entropyInputReseed: decodeHexString(c, "c7e36ae23efe974b763ee423f214555923d66feb9aa174ed"),
		expected:           decodeHexString(c, "1fa1ced6b2f209e430667cc111f2978139360cd4c5267a66d90239a5378149d24783ac4277265f38d440cdacaf479bb734bb5a865b3c67c496d3642337addabd"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_11(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "0aa924d323a39246224f4a6be2b637e206c6cf2cc2872045"),
		nonce:              decodeHexString(c, "0c22c26296517bb99dc92c6d16dc058b"),
		personalization:    decodeHexString(c, "bb1ef0940724b5cf38ddb6cfdca86cf22504620012c27beb65c41cf790294595"),
		entropyInputReseed: decodeHexString(c, "df350b7f1cc53bec3eae292dadb7569deb21eb345c2e1920"),
		expected:           decodeHexString(c, "3e598667b2132e8030ce4cfc92a5b8a750b3e719a101b873a57a17951598528c1aa3a8715b638324ea00a4725ea83f946a249fba6f3620b7f1bf22ff7d092de9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_12(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "38ae861a0dfa5fa4bdb441b1b6891a6a288e53b4ba1128b1"),
		nonce:              decodeHexString(c, "ce4104741c21aec839b336bd063e5b31"),
		personalization:    decodeHexString(c, "5e4a4f9bf0bd36722b456555bfc2044c4fc43ff28ecf7e4e0df3e17ad102c7ca"),
		entropyInputReseed: decodeHexString(c, "4061a0d2b11582fd4bc2a900b5f380f919d5d32af46afa13"),
		expected:           decodeHexString(c, "c85da524498a8b6eddff399864af027d935ec002eaec2c8dc75b1b1d849d461ea6dea80e915a75a37c5c4678074465a9955cb9f07454b6379cb1baf548e92816"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_13(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "b2babc188aba0df48f9ece462bae4c20fa60bbcc76c77edb"),
		nonce:              decodeHexString(c, "39d50f4f3e9c19ca43791213d1e4abea"),
		personalization:    decodeHexString(c, "c01a6242d79f073e112c5ee6bd4a71a1259b2cedd75c5cfb82da2bf1c050e402"),
		entropyInputReseed: decodeHexString(c, "8bbc6d754ad55061252a2b535999f1b580578636927e9327"),
		expected:           decodeHexString(c, "22da3f3c2dcce7f3a3f33129caf2271028424d72bcbd75cc81863b24f96f7501e70bcb5b254cb4d2ca480c2c649d9e8947c2537efedcecd5072d57fd36b75481"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed42_AES192_14(c *C) {
	s.testCTRAfterReseed42_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "ea1e0bd51689f2d121eb55ee5272b6fed7b9ff7dfa8c1f11"),
		nonce:              decodeHexString(c, "31aa1e8f81acaf884048fa056488902a"),
		personalization:    decodeHexString(c, "2ac2e898b00a099612f685473822d2152bcd8f77820a3859b481ef1dfbedabed"),
		entropyInputReseed: decodeHexString(c, "05c7b3fcd858bf5080b7ec08b64c20fd3614d25d9b6f4d1b"),
		expected:           decodeHexString(c, "4397f38e54a15839af1d7f5b190b9f122b1e8534d1d74c53293e9f3727374c2c66ba78488929ee09991598524b14ed9800217d3cb2e02edee7f8410422e36eb5"),
	})
}

func (s *drbgSuite) testCTRAfterReseed44_AES192(c *C, data *testData) {
	s.testCTRAfterReseed(c, 24, data)
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_0(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "b71d3e35697b6a1cf5ca4f0992dae0afd91d6d59fa673440"),
		nonce:              decodeHexString(c, "400b0de03670409c92798535858e5488"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "56e31eb4a1c8bae796207aa520d66161cf06be2d5b46a7aa"),
		expected:           decodeHexString(c, "3dce2aceaaec3fcf59d72ad8cae9267537b0d965b970846d2ca3c52ccc28fe7862ea144074d719e18a7c50a9a188ef289f36d11d20e176c0020d0dbf2c94c282"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_1(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "04a2b8adf83488643d5e6712d09fdb4b0da60e3ebd0e6427"),
		nonce:              decodeHexString(c, "c70ffc3967416d4af2e3f5ef20d7d402"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "b3f248a4ca6f65893e613eaca4e24c99d1ffbb373395a8e8"),
		expected:           decodeHexString(c, "3c7a50f0215cd9b4cf60ea77c06ad209b24df93c9d7a9d288cc4978c18e08835671c9f7a5415de66bbbfaa5c9c2e709998988d86122f985b86fdb502edf1b928"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_2(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "73d68d22b4b6b98cf4ee1a4f6b8f6a035893f03d34eb82a2"),
		nonce:              decodeHexString(c, "76d83338ff6e6571e9597f34a804626d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "dce7c7238d22e63db2edd8f9a48684699f9494b469825166"),
		expected:           decodeHexString(c, "296efd1a98edf4542d34f6938bd8826e525e4661aac5efd02169f552ab12dd0193b6f766b41da0553006cbbaf5b729377eb2c930b71ab8621731a063bb1113b9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_3(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "bee5bfc00f67079585f1159c37e5c2dacaf19a20574b9ae6"),
		nonce:              decodeHexString(c, "8df0c9b5bb7b18df823fe56cf1153d53"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "c1a53cfa7da44eb9420b3eb722daab61858f5efa573f12c1"),
		expected:           decodeHexString(c, "dac250492abf78bf7c433736974f73bb2cb3d3e01cfdf2726d31590863b296451d5b5127507f6d4c50067272faea3aa5790670568948c617af11612496b76ee1"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_4(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "be06744b7cc633132086b2d5a1e9b62c26bde56bd9c0d7d0"),
		nonce:              decodeHexString(c, "705767eeb9b97ce717d7272e75b621b8"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "c9384987518ea2d029ea7c1d41657c55d7793882c8e08ccc"),
		expected:           decodeHexString(c, "ef10c86d134ba16257ad2fa48396134917795724f7ed5257a28788d729ae1393f442bc5f5a5d39884b3b9df56e262be429881a7201695acd21fc2cefdb91ffa7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_5(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "0f06dfb407dc8579617613e5f71654a6adb9b27aef95d8a6"),
		nonce:              decodeHexString(c, "4cd415dd1a32c7e28b8635e54de9287d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "acfa2cbc60e6cd273d649977e567590d05d15e6b5294dd2b"),
		expected:           decodeHexString(c, "93678a280fad7add11776e033588d757540e45144486b3c75d87d2f03ff364ef6dee2c8778d435a4d6975b3cb2530b071b65939703fa017bb52c2c0c3520af20"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_6(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "7e8bef91211ef73daa9eb7d32af79fbe74120b5c183e30a0"),
		nonce:              decodeHexString(c, "e9eb7d6ba6eeb45b67736919de782f15"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "43452ba909077855b2124a5ef90f849ef4f019ce16902a46"),
		expected:           decodeHexString(c, "b1e19cf55da5c1cb0f3646e74bd2df96f6887c779b5474e47faed22e6867dbb2843323c2a329042a155057d242f3fe6f590ae3825138022769bc6e141da51bd9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_7(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "63bc100a116e1482aa5f6d66c18b69f9e3deea25203dd7ae"),
		nonce:              decodeHexString(c, "cf9b2179e3e8d651fddae5ca8b220c0d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "2e7bac4585deea2a93f4069f9896544aefc4a218c141a25e"),
		expected:           decodeHexString(c, "9624977f135c298afffcd60b5be0ff62dd29a2d891389b6b5080e9d88bac609de4c716f06b47c1b7cc92f66cb5f7a4d9da906f084d5d45aee9bbef2624309ada"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_8(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "4378bc78e40b931889ce78510ee527fd96a5751ea77cd877"),
		nonce:              decodeHexString(c, "f4bff2f1da886321eb94f6ba34ef0cb9"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "0f25a2c8d40a040b2bfdae68c567194325c94d8aa57393cb"),
		expected:           decodeHexString(c, "160dbe0b5f8aab91b99ee58224e31eb5e8007d43e0eaee13862a26a1bbaf099f8d93ed4d2112114dc952a0cd8c266f82d5ccd80a5b35f5ffb3b211d0cba51103"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_9(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "8cc5c101ece3cadcce50ad2b805a4c79c51c6795af716680"),
		nonce:              decodeHexString(c, "3a166815aa0c217c0c147eeca8ef1dbb"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "e945678b0333a0b4ded648679d73900942ebe7bb9ad1efaf"),
		expected:           decodeHexString(c, "4ac17acf725fa2c5decf3596cc859bedbc1446746b3940e6d28a3759c798f59162f176b8591b70ac58ea83283f3739ff825fc1c53ef3d91b84832c162a4ceb5c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_10(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "36284ddcea07dcab15d5ca8c6c6ecaa7d1918a383e65a15e"),
		nonce:              decodeHexString(c, "8a20929064fd5f29e698ae1347c7d755"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "de970d724526e068c8a42674bf99905a3e1615bc03631225"),
		expected:           decodeHexString(c, "67e27a41900db00670799b61de771e03742582e8a582256862890d332c858eaee4ed310100bb4323a306b1de9bde3f782219e3e79660444a4cbc44a5c48b07ef"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_11(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "a6281505d016f45227cb86ffb329499a89449d7e7d1f96ae"),
		nonce:              decodeHexString(c, "44ed900e03be20998d0de6a4c8a647ca"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "3ce102767f579dd9e099529ab6480596c38adb22e922c1bf"),
		expected:           decodeHexString(c, "e3be72f6c8cee6ae1702fc106011500872bd43a233587cdd2745f69d12148f33bd56098a706d95fa7531b78332ad76d1b7a7a1e16b6904ed5364b3d1b3230a74"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_12(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "d436c3b46a8f2a04b07ebe213b3d6e464138f8015dafd079"),
		nonce:              decodeHexString(c, "f6dfa2e69075ff6765c4a3ef9d7f8f83"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "195c37c4e2f8c06e40c79cd4bcd55f1d82ec2daefa1a5f83"),
		expected:           decodeHexString(c, "29c9204070d4c255ef0e1b52fba7450ee0e5ae0d268bd8aa9d1738161f22fb95505c1fa6a6b330918e75c4bb8f96fa2b9ac02b5ddb2551bd497e561ae1ec845d"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_13(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "b69f6caea2d9cb3515acd99d2456426b79950594f5d7be45"),
		nonce:              decodeHexString(c, "76d66fbf25f418d8b39a14251917a597"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "fc3a32e809589e3358b798b1b5d642d0f9c1dbe8952f8032"),
		expected:           decodeHexString(c, "7ccfe74c243da01c74381d64e7b2acd37fb4f3d24412e7573abdd9db91e2797013a1d672e909298e5165f29a53096ddd5b90220548b2173c75a0834395d12ead"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed44_AES192_14(c *C) {
	s.testCTRAfterReseed44_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "e8d801574677199bd866197946d2accf3cd1c61730553cbd"),
		nonce:              decodeHexString(c, "afb32d283ca1546dbd5075a903828275"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "34f8c96c52c6c48aa0f39d20502f5b031930667af293bbaa"),
		expected:           decodeHexString(c, "5b8a8cc554a70e64d1e8f080fa20bee126b439147c6c56fdf698e96c5c3e89f20fefed6863669fb8f58ea42e2fcbfa9a40a4fca77a7ebde43956649022c6f475"),
	})
}

func (s *drbgSuite) testCTRAfterReseed46_AES192(c *C, data *testData) {
	s.testCTRAfterReseed(c, 24, data)
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_0(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "f4936a732fdf5668577db79e7f93e91c4cda1d10b5781b5e"),
		nonce:              decodeHexString(c, "c60b0b05600152f033f056dce7af7915"),
		personalization:    decodeHexString(c, "6f0cf012d0e495051fd6c2c96e581925960271b359a705a65075bc5ec16ddc6e"),
		entropyInputReseed: decodeHexString(c, "bf21ff3630b16acef87602badb5ac2d640241eb481a0c8d9"),
		expected:           decodeHexString(c, "18931fb1dc255ddea2e6d4351eb6f329c234e1d4962060aceddf0c9fa99509cbc64c042258f8daf57c120d7a2daca95f40734f59206e6994aa6395cd80a1fd54"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_1(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "3d185c3029d1292e8dc0748afd681c9e8c41077774217b9c"),
		nonce:              decodeHexString(c, "ec68b8b7da238edc37824f55ebff29f8"),
		personalization:    decodeHexString(c, "805cc941c1ddc1a4ae3a1ff9da0145e5c256fd63dfecc5948e0e24a92d8adbcb"),
		entropyInputReseed: decodeHexString(c, "b45d6e410335da94a17841676579de67781ecba257131385"),
		expected:           decodeHexString(c, "fb952f3f623ef2833a4d8113019174b9127f84fa593c380d03ad11643d5a35735848d93009b458ef41643c2666da6deb9432398c4e7ec2480b9c7b879c099fd1"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_2(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "c336f7c8121c4f6df206a94efc5dc130e8ef83dd487ecd56"),
		nonce:              decodeHexString(c, "615b079c738012a3da7fe8e90c4d8ac5"),
		personalization:    decodeHexString(c, "8bca6f65edd74f3c88eb568e2fadcd1c7d3d717abf80a265db3590e0fe50f5f7"),
		entropyInputReseed: decodeHexString(c, "1a02bd0cc158c9186cfcd4cb15328624d37be1e896f0bd0f"),
		expected:           decodeHexString(c, "b446311649e616fdf4088d4f9fd509722b3d9975f9396582258b9892a6cb5f5ab4e261a8156b3522910efc20645428ff3c0f6523284e5c5629f974613d4942b7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_3(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "78ca16d77c8f037592dd30f5e3e0df30cbb8b5c425f843a2"),
		nonce:              decodeHexString(c, "f7e80ed276af5fa1cb967f331f129826"),
		personalization:    decodeHexString(c, "7cf5342f371d56b31e70b6c552aa96eb00009c15a5e0e3eb8f6b9cea4c96eb5f"),
		entropyInputReseed: decodeHexString(c, "4d582b73b54f07e3b5982f12b7e938889a720ed0389c76a1"),
		expected:           decodeHexString(c, "912a9fa85beff2ecd1b8d52780d5a5674e2bf22da5f9f260432691ebbd432c82751968d225bd06c681af89acd2ba483a37214abfdbe4aa8819c64632e3b40408"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_4(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "ff5115ccbc5fc415dd57aba3102d23f6f119aac14dc95d1b"),
		nonce:              decodeHexString(c, "d8a0c6e2cbd0c0b3979d9afaa32362a3"),
		personalization:    decodeHexString(c, "4117457f1256db46943166c8cfe03ccb40a433369f0bd3bb2eeee28dcce44b65"),
		entropyInputReseed: decodeHexString(c, "11bf85b6348bceb863b6a752cb651572a4e750b1d1010780"),
		expected:           decodeHexString(c, "abbba83d02fa0bad5ff6ace6a92aab07b3aa2faac6c5fecb46af4cb43904227fd9f34322b3ac902982b23b55804658bad3f975674051baf554065785ec5bad80"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_5(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "9eb4249f7060d6c3ee574b63f84d4dd7759038d32d15b24b"),
		nonce:              decodeHexString(c, "6491e9d14fa7edd40249fb5592ad8ebc"),
		personalization:    decodeHexString(c, "ca3044f43913be0345a74a998873332ca3623282fd03b00bce0cfd978e0d9b1d"),
		entropyInputReseed: decodeHexString(c, "f658346da825957f1d1429dc32d7251846febe1c3ae6cc3b"),
		expected:           decodeHexString(c, "027f5cbe588dd39c52691e5205039813a574d4543413d092af0d0b7542fb2552a7f583d19600b897cce9e8ac1c744fa499fe0d704e897d76d87aebf4fdc94d19"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_6(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "f418d3a2d3a104a325f82ea17e2b03e3dc471faae783413a"),
		nonce:              decodeHexString(c, "41fddacc56c81807b42e441288da6578"),
		personalization:    decodeHexString(c, "955f188023da79f94f3cfde71edff2ec5a6bab24925f526fc02d820871ce4d9e"),
		entropyInputReseed: decodeHexString(c, "3924cde43683614eb4396fce38d65a498f87e6c374fc2308"),
		expected:           decodeHexString(c, "975cb8c817ac9027e54ab9effcb6f821f4688f9cfcdf5cf9c0f584a5b512236001681bb985d913d5236bed8d536af05eebee87cdc7756af98b2299012640897e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_7(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "b97a78cc2488614dedcc621706ae8096a08a27ed232372ea"),
		nonce:              decodeHexString(c, "685b6c16e3d3f72568644450c811cafa"),
		personalization:    decodeHexString(c, "4ace80d8c0bbb9ef5e1a6bc24a7b7f66fd86c93f95fc7cf14308383c8272ee6a"),
		entropyInputReseed: decodeHexString(c, "9aacfaf811cdcda8655c6af2eefb24caed4a9068d1f6f558"),
		expected:           decodeHexString(c, "40ef287ac7537cf3701b75bc670add6eb941e9358be849d4b7fbdf4c2d4ca94c03c5a801b30a9e9a7441c467183db818fd9eeaca9b978bfc1793c191427da281"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_8(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "f207c8a0c27c3e3f22c6bb8a5bead3576140dd7d8b603246"),
		nonce:              decodeHexString(c, "6480607508afa5a0d4cfc0912fb9f715"),
		personalization:    decodeHexString(c, "541c21c0f85f0f4c58fd04452a7ae4117230ef2fea316ed5e00b74223129156e"),
		entropyInputReseed: decodeHexString(c, "aedf55e5cdcc129e7a6df93463680059a0f73bf91176f69f"),
		expected:           decodeHexString(c, "7f88330fa0e6803922f9c850abe5a83bc1d7c9d5e47d175cc0db2271cd72fce9ccc205df331b8f39858f778f733a4b7092d7805d27ad5f491cb2a34d7982c809"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_9(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "6279a3a52124465e506a00166b318520884625f3d1f5cefd"),
		nonce:              decodeHexString(c, "59a0c8eb4dcf7c54ffaa399dac761365"),
		personalization:    decodeHexString(c, "af79294e87c6df705ec4cbd0f4acc05befc406ae7a52daca1cd4a5f7ec56a989"),
		entropyInputReseed: decodeHexString(c, "5320e0c824b92e9db922f28ded30936bb1a0f671702bdacd"),
		expected:           decodeHexString(c, "152743c5d4a7f9c5726d5d5a9fde8c1e638f2fe6409728f2a6e8a95f307cd2a88e22d10f512cf3ea4b3845396359e181ccfb0272f91bf0795a85639ae5cef014"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_10(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "af2be0eec2c10ba57ff2efeca5e5a6350d6b734bc219f260"),
		nonce:              decodeHexString(c, "99eda95f1bbd4c0793bc34b7fc1e2397"),
		personalization:    decodeHexString(c, "4d616364d52a6f94df53e95c5504f1c5f56d1b4d030dabaedcb4bcf1c81ee5dc"),
		entropyInputReseed: decodeHexString(c, "63efa89965db281caac52c1de62d1f53d99de880e65ab62e"),
		expected:           decodeHexString(c, "e978da1f4149e4309f43213e313e5860e1f1683afa9749a938ae08252bf8489eb30d24723d19e491fcd2e0844b1697dec3823e1fb6fe807c7f66f6e803e7fec3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_11(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "d397a9d80f77caf5becb65b124f3788ebbec7c3adfa6a8bf"),
		nonce:              decodeHexString(c, "060b856fa8db257b163ffa8e1c1ad2ed"),
		personalization:    decodeHexString(c, "0453332dfdfd39a974b95d8491ba57c034762d63dc0a46e65e066c7df6d9ef91"),
		entropyInputReseed: decodeHexString(c, "5bf4455ef30862d8c242ba07b12c7a9cf255e1673c5c3b98"),
		expected:           decodeHexString(c, "4113acc3c155e96da28dc1c5aa369ad5d1bc4a839312d6f79d6e1514696f9a2c5231577dce79e8cc9d949cbbd30339c4e525d017106c4b2f64b548c55374af6c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_12(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "bda88e805570fbbe77b9d53fe5bc443972a0094e54362c6d"),
		nonce:              decodeHexString(c, "83ac35ef07af291d6e46d8a9e4a2f40a"),
		personalization:    decodeHexString(c, "b5ff036854fe4b3f674b1131733b8aaf11eca1dbb2bf6834f3d236e6e0ffaf43"),
		entropyInputReseed: decodeHexString(c, "edd2f0db20f4313ce36dec00a85ff6f70e00d506f8dbad68"),
		expected:           decodeHexString(c, "97b45bc56e89d40d8346e8fe3c3e3c9b230f19256e01c40364090e3081db192230465b8ef13b9d58e2d6218aaadbc4afd6f91b83573e6cf7f057299d75826d66"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_13(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "2880d806b5dca22e2ceb0e242bd0a5d64758f98b7032fa5d"),
		nonce:              decodeHexString(c, "0f4811c8f4942e69be1e7d154c730672"),
		personalization:    decodeHexString(c, "632a70fdf2b9cee1f45a3d18a9672ddff9b22a01b825262e641e7cc6e0c9ca46"),
		entropyInputReseed: decodeHexString(c, "3b4f705e1f4704b219b265e8359c167d1712b83e5600006a"),
		expected:           decodeHexString(c, "2538ebdd9c7392193aa2df2efe8a9ce1c591b1e65e1b8d0837a6c1826dd986001cd169a85d66fc6cb672d4daa3d2ab98c3dcde71047ce73db4b690ea29d320db"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed46_AES192_14(c *C) {
	s.testCTRAfterReseed46_AES192(c, &testData{
		entropyInput:       decodeHexString(c, "2046344904a1d3adbe32476679f6b45ba987eb1f9af7a968"),
		nonce:              decodeHexString(c, "cfbb7761deed18fe83250fd128b7d56a"),
		personalization:    decodeHexString(c, "2030067246e0fc3310b411195c5fba128aa1ec823051081af0ad3a433244b4b1"),
		entropyInputReseed: decodeHexString(c, "fcd07114a56039f5ce591e907b3af1472b0f07c8f9ad082c"),
		expected:           decodeHexString(c, "25fc78ce65a56f4bbc1bf3ad32716ed1bc917cc21f519084c5ff17587a377f2ef7830c1e4f7cf381be260f5cb9397c4ccaea42efc499363edcc14bfbe54120f2"),
	})
}

func (s *drbgSuite) testCTRAfterReseed48_AES256(c *C, data *testData) {
	s.testCTRAfterReseed(c, 32, data)
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_0(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "2d4c9f46b981c6a0b2b5d8c69391e569ff13851437ebc0fc00d616340252fed5"),
		nonce:              decodeHexString(c, "0bf814b411f65ec4866be1abb59d3c32"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "93500fae4fa32b86033b7a7bac9d37e710dcc67ca266bc8607d665937766d207"),
		expected:           decodeHexString(c, "322dd28670e75c0ea638f3cb68d6a9d6e50ddfd052b772a7b1d78263a7b8978b6740c2b65a9550c3a76325866fa97e16d74006bc96f26249b9f0a90d076f08e5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_1(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "200f096b76e3bf2f40133ae6649221084f0afb11f96fe86a4987ae7b1159d032"),
		nonce:              decodeHexString(c, "3be56f6c0ae289dfc636f96cff5daaa1"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "895133f4f2d1be25ec929d42e904dbc7749939ad7022a90360a743fd2c3f483c"),
		expected:           decodeHexString(c, "bf12bf4d8eb6bbbd9f91a2ef48c6bc6524a133dde3c8d4f13d4b5cdae3b9e041b98c8650ada9e1f2b5df01d875470b220cacad0ee887080c271929f695204b66"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_2(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "1cc5a086831fac6ba046b7f56c4ea5ba7bcf9d851b5051254c4683bfed7a26f9"),
		nonce:              decodeHexString(c, "a8d42ca3b08c9c974fa2c2eceb5a71e7"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "e8c174c621af92c5012fc4caca8d1fb72ea7998f5f78a6cd5f3f250f330f0c74"),
		expected:           decodeHexString(c, "6654d831403693591476213bee7bea644c5058f93454e89ea5b348bc5354e2d8abac00d53b3879e2c89bc8f490969e42d738ba37432822df859d631cfc86cd40"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_3(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "6ba5e815274e5cf4b2467743a8333c5c5292329a96f0aea4fdc9a1808b312c62"),
		nonce:              decodeHexString(c, "2abe3c2f11c90ec9b684e1cb3fb0bde6"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "bc7257f625cc1095366d7eddb793ea75ad2c5a475514d53056659423e54cd001"),
		expected:           decodeHexString(c, "b95f8d6258515a67c51f96f8201c0b5445142cde38dab3cff2b527a4e5dca5eee15f79cf073345f3438b1cd507b2fe6ce1569707fe0c288b76bf85e1bf1a0419"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_4(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "14598d23e61d003bf321a2b4816f0a7ea3ef6de1ad6983f93f26b1c1630d588b"),
		nonce:              decodeHexString(c, "2fcefe8c6a93cef35a925eb023179f02"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "42edae478f8ba6d45e97a43906aa2a623ab60403f5f60a4c40548f0dededba4b"),
		expected:           decodeHexString(c, "766ae36c6e9c482c6fa2e7fc1e251dc35b2e2ae645a79c2b8d5c0bd7f520b0f4de1b68419c4dcea07516e255e6cbe96007a25396f93f781b36c9d2ca32361433"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_5(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "b553899082c7835484a2cb1114ceb18fcb26a7b01db8d7cbfcea9c35a64e111f"),
		nonce:              decodeHexString(c, "2e814d7171736aee9a47f994e7639edf"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "53ff45e728979cbb9054dca930da5a54f1c603375621b5c8be0652132f587f0e"),
		expected:           decodeHexString(c, "0693d0a13fb4848dcfb5bfe4a9a02227d3984103ce39bb8c40d7cb224bc9281087d797a5333375052bfc352ea88da1c9368c3e250e095b12091f6b6f12605f46"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_6(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "cb15c90bc72df4a4aded92e9a85f0a23019fbf867b5b027a614a0025f9f3ccfe"),
		nonce:              decodeHexString(c, "3b426df8fc90b5bac1f20e8d32487d1a"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "277098c4c04f2e3f47a461e70258d629fdac97e040f13d4ba015160ad7b537b9"),
		expected:           decodeHexString(c, "75328778fe7a63dce1b7c8cedea9d6a9d767dc81791df0481983abfa2d215ae536bf76b5992a10c4a5cb06858b5a4e3c2d8ba4ba9912aebe960393e81e28aa69"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_7(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "a02de2e53e9b72853511acafa59028c358e8dc4a1c70834d4350658b8999acf9"),
		nonce:              decodeHexString(c, "2da017fbfc2b13f21bda1e70de06744b"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "14e7c1af8760d64c74668dd50950835d9881e040ecd625e0025d8c1363bfd764"),
		expected:           decodeHexString(c, "09e04791c2f9bef5297854065212cf1be44c2a5e28e8f90dc184d4e76c6dd09449859e66f45b7e1f4cb22ae51b8d0c537445b7d438b054ef9c7cc7f5a2ba2e19"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_8(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "c9ced65013ee88a54ee90d95ca6189207c22d7fd93f569ec11bf694243b7aa19"),
		nonce:              decodeHexString(c, "4b3b124b7e7f83a88d83645633d7a86a"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "69c08576b88d957abdcbbf038ecb6db865d12b0b0a7d420b64fdb03a26190828"),
		expected:           decodeHexString(c, "b24af1379b88da5fba9785d8ac5fc9fb53cc3db5c71ad8002a3f0862f48487addcf42ddc193bc9088271073026c33cb1b8efd77203d5e9bcd88394e443dbd573"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_9(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "959ad6bcd9f6b2a107199d9593b7f633ecb030246cc9860a41558834070d0a0b"),
		nonce:              decodeHexString(c, "77841f79562da4e48a665645410e1569"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "213da24906da06ff2b9beb1fe504149636a8acd67001fe326bfabd038a7148f3"),
		expected:           decodeHexString(c, "335748e390ea7c23193cdf672f3182656b9e44e73aff8f38239b0657d8258c2b1d40458a0fe201010b36ede62206ce67c198323b7cd1d81b61aa25a0f5211e95"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_10(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "b9ffca2a28b4b535c2ad53447a2b537c5fd673d2eb2a6e980e8434ec7bec21a2"),
		nonce:              decodeHexString(c, "d23a376451fc7e0a6a0d20159704e9fe"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "27de4e53ba25e74e08a98dc2b96df439fffa0cf211a522c0a92ef1b60830c308"),
		expected:           decodeHexString(c, "ebb300303bf8bcb9771a2fbc755359cc8a8de2d8245bf4acb2b516e2a8bc7191ea477dd84a4c5a19c2c4cd09b8233d58015e4fe9c0f0c601768de0af3f1636ac"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_11(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "4ce24a78795507a537b32c127d949c7df90322a8d5038e259d4cad7d21889e09"),
		nonce:              decodeHexString(c, "1ec7848691ce551876028d24c4d974e0"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "4042584f1c000059c2a1d73c6028567b12d5ef2adac3754f32f41a61ea65fe06"),
		expected:           decodeHexString(c, "b5cbd3ad01d216eb4873ae66244cc6137fa7b46cfea2dd603b4eb7e2ca0a92cfff78c469c4088c623dc2722b187fb8783b4ec10d0c93037dc213d414d936cccc"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_12(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "ce8dafddf08f0321b0f07a825282b4534011786f04288678cbd9f340752a9ac6"),
		nonce:              decodeHexString(c, "d92ae02e9b540b68128419bb628b9074"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "eed6947973735b05dd5468a662802151b30fbde6c956c8f068546c9462cea787"),
		expected:           decodeHexString(c, "496f69fa8565558bfde8b67e990d5f446a7cd668ba0aa10d1eb1710ef64798d7d8c7e08db654409e4c626c0503f3779f14a9b2be22905fbf0c49c30570024953"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_13(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "f3ab5125ec2dbb3dd98e4f0253af3cd23a85f4f0cb01c745f421032b4f0c8633"),
		nonce:              decodeHexString(c, "85204376c77ca3a99a6621354991f05a"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "69167e80478389ce33426502a6f7dd96d31e2cf7864bc8e08caf41a0bcb6e774"),
		expected:           decodeHexString(c, "e6adcd3529afd0557c1951b63256c6b7b423b12710b5f4f87715a8ff2156c07cbea53f29a67c60b010dc4c457504dd8ae4ae3f92dab3c2c46310f4616290cab0"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed48_AES256_14(c *C) {
	s.testCTRAfterReseed48_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "67de0f88bd02179381c03be6295adba3c102f5ee74f85a96eebead925d0e80e0"),
		nonce:              decodeHexString(c, "9ec1ef1fe9ee308ea9c4d2447b9eabea"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "1251331a10f9fbe938485858352470c58c4729a9d9c47c645d0626152ddb2121"),
		expected:           decodeHexString(c, "d669b7d6dc83b16e2f8191d216ab0be3523981b4cca4020d589f4d79b8926838334fbb7ef48265daa1091ef285fec2786c81e71be4392c8244e436598d0af391"),
	})
}

func (s *drbgSuite) testCTRAfterReseed50_AES256(c *C, data *testData) {
	s.testCTRAfterReseed(c, 32, data)
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_0(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "5bb14bec3a2e435acab8b891f075107df387902cb2cd996021b1a1245d4ea2b5"),
		nonce:              decodeHexString(c, "12ac7f444e247f770d2f4d0a65fdab4e"),
		personalization:    decodeHexString(c, "2e957d53cba5a6b9b8a2ce4369bb885c0931788015b9fe5ac3c01a7ec5eacd70"),
		entropyInputReseed: decodeHexString(c, "19f30c84f6dbf1caf68cbec3d4bb90e5e8f5716eae8c1bbadaba99a2a2bd4eb2"),
		expected:           decodeHexString(c, "b7dd8ac2c5eaa97c779fe46cc793b9b1e7b940c318d3b531744b42856f298264e45f9a0aca5da93e7f34f0ebc0ed0ea32c009e3e03cf01320c9a839807575405"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_1(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "5e1a564a70f593c1c0b07c9906455bd9f5ce7ad92eb344a9cceb12f5576d7d9c"),
		nonce:              decodeHexString(c, "45e093e587341f6cb8f3deffddc4dc4d"),
		personalization:    decodeHexString(c, "b61714ba7ed339a24635c0bd4f4db496b74631ebbcd14f648de71bd6d7c197ff"),
		entropyInputReseed: decodeHexString(c, "4fcf7ab9daa808ae81eaf728dc74bdf4c123a1e2444e5118c8040142fea50a0b"),
		expected:           decodeHexString(c, "4d56fa065a3b98f9ce21701c00c833bcd439276fc70aaa14185b39f34d80232565c992e2f0fbd9519175751b4057c21ea69d4c553e30e3dc5533d4abd97ab19f"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_2(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "c32238773de8dfdf3bc319a64631c3caf67ab0716e8946eee2fff1fdda96d2ff"),
		nonce:              decodeHexString(c, "ae2b3a16b031c784b80b94b45c8cfaea"),
		personalization:    decodeHexString(c, "b29400e49e0fe24c6418c4da38417f857d53ed61070d467e34049f613568978f"),
		entropyInputReseed: decodeHexString(c, "91c36b0c87587b663583f636a26303f308b7a5dc235cb18086d4e350bd3fb631"),
		expected:           decodeHexString(c, "a1d5a059e6f3c25a1b10613efbfc483095cc257fd98ed2914379bcd8a2ffca2b3d745c32dffdb721ae7a9dea85e0b7a993dbdfec01acaf1097dd9f52ee223a0d"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_3(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "ce80e5656090e097bafc210370213d46f358f77903fcdfb877a0e57f453b4f7a"),
		nonce:              decodeHexString(c, "4515c86448eda28ee63817f36a282ba3"),
		personalization:    decodeHexString(c, "c7875ccf1e5ef1f6d7594296024a71caca6cf53cc86e4e02f86fbb03506fa9a8"),
		entropyInputReseed: decodeHexString(c, "8ce6f56cd5b26de59e01ea11509a23e598aff809dfe07df7e4994c99885eb94f"),
		expected:           decodeHexString(c, "41cc565ec349c978bf7c4af28a6ca9b1a59924b23a581a7f3b43ae089690d6ac262c024fc16d56d1b436c8004522f87f5e8ec3851903ea1ec874505a206d1659"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_4(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "417b1a5aa4694acc25ae2fb18ebee5055d691f8908888e608862c831b9936eae"),
		nonce:              decodeHexString(c, "53a227b0468602f6d5ed623b6b552f48"),
		personalization:    decodeHexString(c, "ecbe55cde21a7d74f03408e5fc8b4c162ee06651552fd32a6d40e06c667f95e2"),
		entropyInputReseed: decodeHexString(c, "d1a00e5bf56519c127a17ffca848a2276b02604eb01b9283de5857fa8d19b437"),
		expected:           decodeHexString(c, "ad11375cd7db354fd67302d7065c9ef36dea373f744114ceafeafe6b91479837ec6fd9cdfc29220e84608fb8c1a59bde7022a8f1e31bef034895cf06a8085188"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_5(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "f7f9bc798994317ecaaf3054af3f65494aeb2a235a6e7668afefc4317757abbf"),
		nonce:              decodeHexString(c, "5d9789c2774b8586dcbad413460b7cb1"),
		personalization:    decodeHexString(c, "8c188fe310bd4200bf84b57617ac0daf2c373ab21df7b0e561aabbd2e3ac19ef"),
		entropyInputReseed: decodeHexString(c, "ed53ec2bd6ed5458a5762c38b5c59282f6e5565c3babdde661bf602a33d6f08d"),
		expected:           decodeHexString(c, "27e7cbebd67c9d82bc5e796710b570e499e0bf9ba39054bb0c989a045b275f5f0c089e5a01ec0bb74cf29e553dc2b52c0b53a3037b6292a413299c9d03aedff3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_6(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "3a670102a446db0567742d42eecda30469c96211f8e7fdf8bb7201cc5e602481"),
		nonce:              decodeHexString(c, "9bf138ee6af50a1bc22749da1f36e6fe"),
		personalization:    decodeHexString(c, "16b8e84e249eeb2d26f89f4797f3ff38a068718cc03d14c6556c255e1cc6f66d"),
		entropyInputReseed: decodeHexString(c, "13d8160e0670aca840d95e0c396115192ff8418cfa459734b6e35c4a4144efb1"),
		expected:           decodeHexString(c, "5ad8d437d21a11c37f9e950aab0e741b7ba1798a9fb8eb166d40eec42f9c07d272fe7d95b155611fc6e5a45d9e355a55261a28db17eaad373c46b4eff6a14b59"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_7(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "05eee5cf3a148a84f14dbe86cbb0104e40893bb0b4a712247b8dd52e4a66ccb9"),
		nonce:              decodeHexString(c, "e1e5c4830fd73e87e6346c55e216d075"),
		personalization:    decodeHexString(c, "bc41aafbcc7e63c02d7e9c3fb95518b0188867567c65735c12f13f5ab90e788b"),
		entropyInputReseed: decodeHexString(c, "702a6a0588e72b9c952743645e3d00b35a0c8b0c2c39da09a2e43e91b4dacb6d"),
		expected:           decodeHexString(c, "f7de81c26c2f78b42c336a8e0cddde2581d4d06d4090750eff3e43816f6ea33f56beab6f78793ac45dd4bc0a1d34f49060f72fab0f8f31ac5b7e980e346e2f93"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_8(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "6c2aaeac3012fc4acc8d35c671f5d88fa25f50d8c80c031ac5e894220bcf6fbf"),
		nonce:              decodeHexString(c, "baac5cc170847c815a76fe6e7f9a3da8"),
		personalization:    decodeHexString(c, "8db29b7ca6684a13ede4025f6000482a379f745604a7d5bcbf60a48ef6cd8db2"),
		entropyInputReseed: decodeHexString(c, "64e9862f9e663661b32a8e27a70b2a3c0ecd3f1ca3c6e199995b1b587ba31e0c"),
		expected:           decodeHexString(c, "fa74549270c648472263e0a79efb8239f0369679cd461fc68734f10432cd266b5bd2df0b50cd307bf479ac63d5d33dd65017ad51b8b8577eb42a45acad373fc7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_9(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "23881618de81ab18a1e31596ae03632a500ee8d751c4bd30972277e3abddb48d"),
		nonce:              decodeHexString(c, "88cd130a12f92aad96e16b13dadcd9dd"),
		personalization:    decodeHexString(c, "2d9dae1dcd0b7b57108880c322514165240140d875f2fc829d9b2ef99dd371c8"),
		entropyInputReseed: decodeHexString(c, "8575f16ac42dce0de11323905354991f1b2e85d75c2c89302f5a634cb0da2437"),
		expected:           decodeHexString(c, "66308b40e12dcb286839f24d88cd19eb46c4490dcafa92d8ea19d0b26f73e15150e92c9e7918a2f18c9b26599c9f19a813b4f01ed566174127feaefc5d151ff4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_10(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "cbde0b364db22d5107fcb29b0662847015062fbe180f9dd13f8b6a0fa79ce7db"),
		nonce:              decodeHexString(c, "f1294dd5526d94972eb08fb3fab783ff"),
		personalization:    decodeHexString(c, "7b1d46976d6d18f0ad0c39286b9a9d5549c6aaabdf1df0f0285d2eece4a29a58"),
		entropyInputReseed: decodeHexString(c, "3d71f3c4f5eae778333e65315664d44d3a0a58865bddfd62d22f019dcf2bcbdb"),
		expected:           decodeHexString(c, "56f71f0d48804e0f2eac77f5d34f7bdc5e73b4e6421d30623a50860a4efb449b4bdab3918ba94a898d013f1513a40145067310744e9a4198c5d3150fbdcab5ba"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_11(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "8bc68fd8e3e4254dd1cc178cad2271961967331f3a9bf3a4b440407ff0cd5747"),
		nonce:              decodeHexString(c, "f6d92f1633a1c415cba8d13597965f4d"),
		personalization:    decodeHexString(c, "7f5de45bd123b5f835071d51be22e512c86690df17ac9d2109ddf8e2d7d4a65e"),
		entropyInputReseed: decodeHexString(c, "2203afda11d39aca507939b0cdc1b71a46ec50c8fc75cad87e8664c143913d07"),
		expected:           decodeHexString(c, "5e921322aaf8030122a6814c9e33a2b67c02056eafd7fca457dfbdf5527d3ef7bb9505d969dc353155c7c9234caa5004c3fa6c8e6380b9e25cd6c2c36c840fc6"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_12(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "22e2db91efbe30b53fa643d89e607a1b7eeb1171caf9a50af5ba5d8610bec9b2"),
		nonce:              decodeHexString(c, "7e7d51f89c10aea9c13ad03a17a6f208"),
		personalization:    decodeHexString(c, "8a7bc17552a552db2d6c96bdfe93f4ed61f1b11bf9f6903b4fe306638fe0357f"),
		entropyInputReseed: decodeHexString(c, "ec219ccf1f5655a2481c6af35d8866f354472bf25744731141bef7463687fd28"),
		expected:           decodeHexString(c, "19c42f82f8ffba0db3587dbddacb95376be4ef5546f33124ffc34da499bbdcb15a17727b5f414d010c22728e8f9c721ea0e0ba5dc68f7b29247bfd04946b9dad"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_13(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "4f5673ce798b07ee691b0c426d529eb6c938f16ff330472fc6f60680a3549fd3"),
		nonce:              decodeHexString(c, "a07df7d8762412dc61a9d78ba0244d5d"),
		personalization:    decodeHexString(c, "9fdcb17da44192caad6b570dd5e75be66c3b303ca7c14bf720c94a2def34ddc3"),
		entropyInputReseed: decodeHexString(c, "4548efd4fdc06df54580f1426e1be1455f1e6d724b07480974a4c6f16b16a190"),
		expected:           decodeHexString(c, "a172fdf2cd1ad46da5a90c00fe392bbb5b3b4405a077108a1949b54c052364ebdcdad34eb9eac93ff91e5e13cc67f084331021f8db723b46fcdc1378157a6d0a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed50_AES256_14(c *C) {
	s.testCTRAfterReseed50_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "abc9f9d53810de8e38bad119d5234017c66ecbd41021861fa28256e73d3f701b"),
		nonce:              decodeHexString(c, "194d4d4c8e64bdd96cab79e23d2126e8"),
		personalization:    decodeHexString(c, "21dc8141c892ea173637525753c11f1158fe74975ee55ffe76c8a439a369fd25"),
		entropyInputReseed: decodeHexString(c, "e999c9d8b6ecae35a4e0741eb944123b9bfb82424dcae184ee36bab4cedd5470"),
		expected:           decodeHexString(c, "30c328b6f8cd1ed86d106d40b724f942bdbcd903811f4b8c9dd0d2546638750e51427ecdb517a916f8ae11900c4ad73db1bd1f235cf8cef81c60c75cfc4ee323"),
	})
}

func (s *drbgSuite) testCTRAfterReseed52_AES256(c *C, data *testData) {
	s.testCTRAfterReseed(c, 32, data)
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_0(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "a89d08185b539a830b1e9b74c01f59e2b75bd2e2cbcf95c185a83a8069439e42"),
		nonce:              decodeHexString(c, "c675e3b634b075db09789e5d8a39c5e8"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "0ed8e63b823af5476dcb9702daf46185d3f4953df704749d3dea2fbe0c7a46dd"),
		expected:           decodeHexString(c, "61f1fb64c0668747d270d4fab17c34db3a69829ea08fe43ec359ae174ffb0caae8bcba3a4fffb5b29b900f0e2ef2394c39292bf295623f894617ce9500228bb4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_1(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "00c312cba2ec5d72f9549e2a1414c973f4e9ed70407971f58ccbcc85720f1fa5"),
		nonce:              decodeHexString(c, "031e82c60be96498705e6dabf4c550b7"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "084b11ecaefe51dbb7a2651f45b0e181928c65cec575f7630dbf9f49c084a584"),
		expected:           decodeHexString(c, "eb2c76ed3e9467ecf9fa642b872cbdf340a2e1f7116f5ba59eccef7be82765620fa3507a3f870bfc8574041dbb9e7b8a0db6906bdee0bc5dc144922d670ceed4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_2(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "42cf0a3b9f081f46945c37822c4cfa65cb6fb624fbc56fd7120c159fc5585283"),
		nonce:              decodeHexString(c, "96e4b7f661f0e1aa7e3561d06bac1430"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "293e309dbc4b90f805ad2e7dd406291002c28384cb29bfc72c305a93db6c502a"),
		expected:           decodeHexString(c, "9485208c002e4e27f80bdfed3c1bf327e3c0f4f074fa8f60eed40752c288c5398a77643dd9a7ed508100b047b82d429f3b1806f050e0ad57f97141bb7a5d99c7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_3(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "4d53cca2565779f6cf962367bb3793b0fca3feafee09dfd7d3b4d9bf0ba5aafd"),
		nonce:              decodeHexString(c, "9a51814c357ee87441fe027760931033"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "da0de5a7a54dc3a6c874d8e5b31c7cd2c6d2b58344321ecfb1f98d42807d6447"),
		expected:           decodeHexString(c, "7274b227d024475d5248cbf56791c9bef918e25e28659e6bcc7d0450e9c25b81c5b6442661d59f972ee9594528979a0d92c14dc93f4adddb03ea48b15dc61cf3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_4(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "1597c35f95f94f12bb94a1a47a0696f468a8725a6793d4d9848aa06f2ca08682"),
		nonce:              decodeHexString(c, "44dd56839ea193e5a1fc34e9c611756b"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "ae7e1793dbfec60862c0bc91293d6922159313084810cc5069b75df1cb87832a"),
		expected:           decodeHexString(c, "4f7ca39c8e906d126fdcebaa89a28ceb638b3dd5b9a2af0e2708b4bc5ffb8c28eba3d42b3bc7498e4cd371672049dd9b83472e1e47b98df77f15d144ada6788d"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_5(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "b87daa167294e273ead3150928c7583cf808f334adbe8c56b181fcf0325d8fc5"),
		nonce:              decodeHexString(c, "98c039bc4218a3cd763e40b7b65e8aa5"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "7cd899b6d3762fa4ce273b81114b085d6f108cecd01e7606b64046807e6344e6"),
		expected:           decodeHexString(c, "68c3a61438c00096c15917e7941fda04945ec549479142e84c7f29a1476c37207ced72f8600c1c64613c30a9165781a2d2ef17606cd5cdb6fe590a2cbf992243"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_6(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "8161eb935ea90cccecfced72a10d41eebdc75e5b1ecc1f0d8a08326635d05f11"),
		nonce:              decodeHexString(c, "1f9cdac6aec9e74272f40a5287488978"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "d13414ac422e7c359703065100d06e64c71daa4998e65ba4ca7170b31418815e"),
		expected:           decodeHexString(c, "63a31cdbcd3d10f9a3667fd57a281df523ac6fdfdca93f3aa57b0471622401e203d5b0f2846e5eae9ad01ef6ec8c5b6cb0afa1bd244806d0630b1a2342f36054"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_7(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "b8fde0b3bc4608477829f22ef3ec37e665e6ea7535fadbbc6591fcae02431feb"),
		nonce:              decodeHexString(c, "20d596ab902a880476032416b2e80c35"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "47fb3379e4f0d46fe82faf1acfe055a57f63f91870c13fbe16c40de41368477c"),
		expected:           decodeHexString(c, "fb59eaf7e23b7def451f21a3e1a7dc02a48dfb2909332ae949d717d1264f86e9cf9ac476a15679259174d4a77b50525e030345fb9b04a7101ead5f8bd755749c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_8(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "b772f663ad91c0f72f835bb0cdc9ab22a390c057500d2cdcdd0b29f9abcd01d7"),
		nonce:              decodeHexString(c, "9731681cf560d60c2b9786a6618995b9"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "2c93cebe266c4891220f490179b040e41d4174248c900f2dd2bb32fbce0435c6"),
		expected:           decodeHexString(c, "ef87f820566cd82b664c3d40a1186ac80513535c3a2b6e258f6a764dd7b292a017ecbb9d7bdf3409998ae6b3bc31c1e4d4eb876b6b0c5ceb9704e957493572c0"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_9(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "7e313c2cab1c49ea71412236055988ea958a29f1c66ead5daf91ff47cc5e8436"),
		nonce:              decodeHexString(c, "c2b8520efba1ca9785d19ee058cf23e0"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "2a375183426c044e84d7163c0674df324889c3bea2baf057ea93a47ef775a8bf"),
		expected:           decodeHexString(c, "5e6cc25cd2a20b89a8fe894f2f1e726b665441f73c6e45eb41af9901ca6ae62e63e082ef49a1bdc9d113e99abff748467add4c6905b88c4d2c2586733f4b33dd"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_10(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "b27e9c0fb494c09e2a960a5f03491a461cc3304c92ebede9e3ccc748f502a8bb"),
		nonce:              decodeHexString(c, "d0b6a2940d436f09e0e1bd903cc4463d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "f10b9428d0d6009c8a6da2483a147246e20210a3ae82789e9e32d3d8ac5c4f87"),
		expected:           decodeHexString(c, "d0b5946a21551d8408ea54ebcde893ebbbd5ad9eaabdddd2e7601fcecec9a7939182cadbe2ef7ba70bee966a22454549e9d5c13444e442addad8ba4e55f5d749"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_11(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "9875fef6d8b0699145ce20387ed09ecdfcba5dd9bde9267de3a55e038fc64ac7"),
		nonce:              decodeHexString(c, "f3e744f438717b812d02ed4596410ddf"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "a6a1115f2e8b8df21cad820bfc5fe3077bdc6bf88abe07c52f1fbc4c353c2237"),
		expected:           decodeHexString(c, "8908935dc530284984aa1c55def56d7b07c740aa1b208646180f9080f9bac3ced4b9cce30c4a820984c69591c97de703d062df19e211dc203406e188f2122a65"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_12(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "1e8f79f30416f808ec317e40b15dcd935e10e2914b9d83413e2185d0099ffed6"),
		nonce:              decodeHexString(c, "24b6758f5e31b325dc736e6cfb2c36e4"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "20d59d0bbff22f2e991a8b45cdab525dbcae36919193233ca9b08bf10d41f6d1"),
		expected:           decodeHexString(c, "cef29871306da0afa00f6d6343057c2d8993ebe3e0abe0f5a8caafbd672a63862fe6bad69453121ad8a757ea4e5c482ca14729b6e2ce01b3996d4d34a1d832b1"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_13(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "aadc2cce3acf398c8c11bf8e205d6157f9903e8f195246a7bb810adfab0a7628"),
		nonce:              decodeHexString(c, "8f66f880d0796f8ee55545a2268c4652"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "2502ca2af97079a387dfdf2547217f84c3932fcba49177c2a95281f3a289e83e"),
		expected:           decodeHexString(c, "d57dddf87243de2db9a5bae495cc20ebe819e6a4674606199aaab3b55931c1bfa1d133357815b394ff9e810b8373daecd859269871eaf6f56be4a743b1c1997a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed52_AES256_14(c *C) {
	s.testCTRAfterReseed52_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "4eeb688f8aa860047496421617266abee3eef3f88682a79251116bf78016a8f1"),
		nonce:              decodeHexString(c, "bc38d83e891815597c408b0a50a2948d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "aee4a6655f817412f27f4ca686f1476ff38be06abc2dfafed950fd46df03865e"),
		expected:           decodeHexString(c, "2dbec5648c608da2f195f86a41d26a9887a7f75f38ed8d5dc51b8ce67edb10e968c1b054a78d1298cc3e6d8ad361086a0fee9d24cb36fc8434bbaa1442e28287"),
	})
}

func (s *drbgSuite) testCTRAfterReseed54_AES256(c *C, data *testData) {
	s.testCTRAfterReseed(c, 32, data)
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_0(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "4cfb218673346d9d50c922e49b0dfcd090adf04f5c3ba47327dfcd6fa63a785c"),
		nonce:              decodeHexString(c, "016962a7fd2787a24bf6be47ef3783f1"),
		personalization:    decodeHexString(c, "88eeb8e0e83bf3294bdacd6099ebe4bf55ecd9113f71e5ebcb4575f3d6a68a6b"),
		entropyInputReseed: decodeHexString(c, "b7ec46072363834a1b0133f2c23891db4f11a68651f23e3a8b1fdc03b192c7e7"),
		expected:           decodeHexString(c, "a55180a190bef3adaf28f6b795e9f1f3d6dfa1b27dd0467b0c75f5fa931e971475b27cae03a29654e2f40966ea33643040d1400fe677873af8097c1fe9f00298"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_1(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "29cea31e473208a552ad826d25503ebc065d887ddaa83ef9cff83044f2e49bc0"),
		nonce:              decodeHexString(c, "454c1c318f74b332c898f02e951f4fc5"),
		personalization:    decodeHexString(c, "678daeda93305c64c0fd056c9ef42695f40e5af6130821b4a4d706e7013fc523"),
		entropyInputReseed: decodeHexString(c, "2342d3d62acb6d402af757359631b53029ed18d97ef7d6ae9cf7ffc340202808"),
		expected:           decodeHexString(c, "651467aca6454e175f857924e1483294c7bfd3bc2263a1dee903b7eb9bb0899503bf61ec2a9db58e69aac09ac44631e4c7d4c05dc704198706eae2d1a1ef766e"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_2(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "239ea14c16900173fbed0806a3465df483ce981606d9a36880d1ca8db24fc298"),
		nonce:              decodeHexString(c, "3af404ff3262200c22b646ba80bbf538"),
		personalization:    decodeHexString(c, "635737220106b084c641bba005731febb6eae458f0fe38777b2f85b049a171b7"),
		entropyInputReseed: decodeHexString(c, "34519e5f5a23700d3b62cb3f0f362214a88742cc5d112d474f8cfd81a93ace1f"),
		expected:           decodeHexString(c, "d75542ca926444d0ab13d42097fab594c50233e21b5d4639e32c5bc204d3fbe78b583494692e720b0714b5dd647f5ebbba76f1e27028b979c2de7b62f7578768"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_3(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "d8ff66e0e9c26a7985dade71e9f61ba4353b887a09fbc89d77fa9dc739ffc7f8"),
		nonce:              decodeHexString(c, "4ae30b047f6741393e8d7725992c5c44"),
		personalization:    decodeHexString(c, "517e7d941379d25c82c129c10f3ee4dd7eafad1753d7383eaf819702ea93f1ea"),
		entropyInputReseed: decodeHexString(c, "b088ea2cc930d1677fc69d9e605947c598ff674b52742fc6db01775a62d257fd"),
		expected:           decodeHexString(c, "5044f68a7a7b26cfedc06378ba9ea16d47152542934564bcee627824f5b72b595ef3c3d8fdbaeb296c8e1066401ff438d3b3d1d25aecf779034323a2605f9ea8"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_4(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "9173c44abaf926ae00b771bd72c497cd583d8b3c116f32044d6ace54f29af59a"),
		nonce:              decodeHexString(c, "726dabbe474651da7606b65a2bbe0a6f"),
		personalization:    decodeHexString(c, "7a66dd4b42f90a05575cab4608c94d69e74c968d697f66a2ead40d4dc0d53efa"),
		entropyInputReseed: decodeHexString(c, "09f2294f43b68a992509dcfaaf82b30ec473667be779f22b0353d901d21a7047"),
		expected:           decodeHexString(c, "f36d59c8e328ba45b15074bc596962ece0484efc7335932d8d492ecde2552c6df3b52da8baa05dd418cb39b29f8468bde9e882bc11e07a037eccd2047c0b32ae"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_5(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "de5a2a512931c0719332ceb514605f89b305cae62624e6f8b4e59844c461f2bc"),
		nonce:              decodeHexString(c, "742815b8fe399e5f2df45811f654c60e"),
		personalization:    decodeHexString(c, "f869d930288961e43e112ec026deaf76cf5d0012c245eaec571b30c13bb534c4"),
		entropyInputReseed: decodeHexString(c, "aa1c493e8657ab3dc2d778b5845c1610a6d07971e4366666d246c7aa15578b01"),
		expected:           decodeHexString(c, "7df6ede450facd51ddb931f7a817b6c1ff27a3094cf7dd4e25c390bed838ad47b8c03de0a6bcbad37b0d1cb55aab58f6f0357187b2ec22d9e88aa980b6e54d75"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_6(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "77531f8ad007aad31fc105e0eec024d502cf76fd8faacd8b46eb834dfcf8d5ae"),
		nonce:              decodeHexString(c, "37de8baa4b96689793ae6ad99ad3445c"),
		personalization:    decodeHexString(c, "67dda2db559ebc638e182cc5290ccc1bbfc7017af2da6b998b85120529618742"),
		entropyInputReseed: decodeHexString(c, "6cedc868a200edcafc34dbff2bb4bc7851aa08a9f9238b3f2b31a04d66ab5767"),
		expected:           decodeHexString(c, "0feb6cb4bd7774913d1752ec477a43e4cfc1147e8264daa33d907b5f3c2de74460bc7d45d3f174bb7b241256aef2461931b35160f793e98640b4e107e3585dcf"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_7(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "eef297f88d13ed4ca5fce56acb436c3877d7b94d0adb90a37744397e9e846847"),
		nonce:              decodeHexString(c, "6431677c9b85220d1c6b1f786419facd"),
		personalization:    decodeHexString(c, "1476d4b916a8694a45fcd0089f3b6152ed6e92064b1f6b6fb0a313c7aa8efada"),
		entropyInputReseed: decodeHexString(c, "93825a2828662690424b8c6cf8ddbe9cd14b14af8d91984b6676fa6a9242845e"),
		expected:           decodeHexString(c, "a45fbb996a1c35e7c672b16869023c7d1ce81a1e107a4607d2f756f7904526b729858515553e39a7c7f44912a27d8fc7fc61120a6362543398a2b58ccd7a67d3"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_8(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "6f7768480701da57ffb6f65fde52b3076d0d54df325a815ba0089cf966766e69"),
		nonce:              decodeHexString(c, "7f979d876def96d803b1d211173ce499"),
		personalization:    decodeHexString(c, "a48d8cc12457ade11515ec9ddba1274d05a6b34070f04ee427cbd26afc2edf3b"),
		entropyInputReseed: decodeHexString(c, "3dfe53a61bb795537c65fed8ffe09c3f8bb62ffa5e9e26cb2907401c4b8dcc1c"),
		expected:           decodeHexString(c, "e5eb353cf0adbef2e6c62d745876a835659a3a94cdd2328bcca6abf96ad9637be2ff68e27b8e7cc45a3b79d2573661819ec684eab34aba07c1fae6ab81c988cd"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_9(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "2c077bd78a8867c627c856d8f04d7d6f1d2162232b33916a946997f4fb0fcf55"),
		nonce:              decodeHexString(c, "f01236e45755d721d575c4e9304170b0"),
		personalization:    decodeHexString(c, "57efc60e693151fbaf6051de84fa0429b4eaff35feb7e824c2fbbd692fb8d68f"),
		entropyInputReseed: decodeHexString(c, "151febacbf949e1293710a1250d229ec02ba4df10b50abc2742f083e3f923abf"),
		expected:           decodeHexString(c, "a6cc431fb626a926210ff7d3084e133dbf00b22d9877a07c82221b1a5ebd77ea671950160eb298184afc623731a2225b6c67104b85b91022ad9d33e8495acfad"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_10(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "685de1f040678f2e86111f7abf2f7596493baac932dd9ed01fb70aefab40fe03"),
		nonce:              decodeHexString(c, "5ccde240d7ff2293091a58de55602ca9"),
		personalization:    decodeHexString(c, "3b18892feab083b530371d6eb599828a58a76a346f2a25a412f5d4606f2f0baa"),
		entropyInputReseed: decodeHexString(c, "3e7a1cb8bcc4f2c2c6262b1b8fa7bdc20aac98e3f425f781c7d685bb43fe383b"),
		expected:           decodeHexString(c, "d87b0ef23e09b1c6f1267268029528c76b3e3a6648c674fe92486869a47f7892e5660f885d0fd2e6b2a2288561d07575c6606899a6551c4f3e2f14ca75c435c5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_11(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "eb4b9c2e641880e857a2baebc8cbf91824ba35d26399a7caedc11ea42d7c858d"),
		nonce:              decodeHexString(c, "f4819053ac388f6344e8d06fc94067a3"),
		personalization:    decodeHexString(c, "10be74844506566a1db552932c9affb314b89c83434307b873a0126f4cf28a77"),
		entropyInputReseed: decodeHexString(c, "66826d4b8ec8038f7199d96c8495961a0a74e391bd2899f0458059ea4d2edde3"),
		expected:           decodeHexString(c, "83b738bdb863cc7712eabce93fb935a0c01834baa118ae99a6163b5dc05a71b02b93e8ebfd6a20deb3d54e1850f82d96afbab2c13b1faa27c5bb01281802e2f1"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_12(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "6ee0c875d08ce12cb10ffa824be2b75713901dfad2d94309a1ecc4b51adc37cb"),
		nonce:              decodeHexString(c, "c03e8ef10c253943a05801b7c0d04e70"),
		personalization:    decodeHexString(c, "0e717db96ebcf894c0182807eb491a8719cfd433ec02dadab7ba2fbdeab1085c"),
		entropyInputReseed: decodeHexString(c, "cd6d2a324cd38f3aae6ef8e93de701f0725c3c08f07d3570d6c8c01b6525c18c"),
		expected:           decodeHexString(c, "dc5e7aa22a722c62d68391e1a59793ab4f27ef9f1cb2c3247bbf94c339176ae81010c30c75572971be8f78a6cf8cb4c3ff13bdf00c0e3a259ea70306bc0b4b02"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_13(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "600bab60158002dbae0865775713cb02453fb525d5cdaebd8ad32303ab9cf86b"),
		nonce:              decodeHexString(c, "661a0b680f2b7682c157e01c99d83fa9"),
		personalization:    decodeHexString(c, "1c16a811080be74b862f9f64d4ba0ac8964439bda0e560584b7fe8e5d67e9d62"),
		entropyInputReseed: decodeHexString(c, "aa15f26c205d46c20fbcc5e8883df0a2b91ac3139a3c1fb58f1fff17b1ce0d95"),
		expected:           decodeHexString(c, "ab59d0636045cc3463478450df2e6e9e08c209041544fc15d32dedb641756f15207ac7a4dd56dc006ef9e5205466904a47b5e511667c7d141e1b9ae9d9d6a861"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed54_AES256_14(c *C) {
	s.testCTRAfterReseed54_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "ffec8a7f170fff952668728c9c9390e71ba4138fde684053351376d7ab54864e"),
		nonce:              decodeHexString(c, "a27070b710e5aee5fad9c605df1c34d5"),
		personalization:    decodeHexString(c, "dbc5113ef0d4c4e61a274cf6661819fc41bf2f91f2e44f39df43a76cce0c84de"),
		entropyInputReseed: decodeHexString(c, "ede152730475080f4b8cafdd33efd0e7d03529c06834fd5c62cf0708dc961d31"),
		expected:           decodeHexString(c, "0c5098b7bbc8b8e2045dd6ac82508f836c9e059c070e0499bcbe58b20d9843f258a6cbebdc0554686cc04507dc589caff460f0e9b8dbd9d9e6a84a36549e77eb"),
	})
}

func (s *drbgSuite) testCTRAfterReseed56_AES256(c *C, data *testData) {
	s.testCTRAfterReseed(c, 32, data)
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_0(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "830bdfd33486f26f4af9f2a699db1e49652635aed6984e04a0cea2c9a87e43d2"),
		nonce:              decodeHexString(c, "21ede5be36404c34b1b85c2d2369bf09"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "8c721957a6300794862a004574f98af9bbc074ecdde22becb081f360535f3f1f"),
		expected:           decodeHexString(c, "3f63eb5de3a13a3097e25399c3d9ed7d5e6591931461a851ba645bcffdd0c07f2b71cfbb8329bb1934971d1403dc68cafb0bd6ca4e4a6c28976ad5e8bb13a35f"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_1(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "068ce29e91fa6ebe9d39b01e288fbb5c64d5306eeae703d3b74dcdcd64757d8f"),
		nonce:              decodeHexString(c, "c96064d619d4ee605deb0cac78029e0c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "a5f0c736bac2f1e7c7554f51e87279abf01d39213f20e310ab45d0e0262270fd"),
		expected:           decodeHexString(c, "241c13c5f180e17382b03229cb6037a2238e658b0bc7927342833ef0b4511bf80d8d04042a7114485b6aec347da89c64ea5f7d80e8f4abb4b054f2f07ac6e2ee"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_2(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "f22cf7cff5c8f25c3b15d9e64b728ee8d15cc90637e27b64c4643e46e19afb76"),
		nonce:              decodeHexString(c, "aef366b3955f78f1cc43ee008fc88b7d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "17c1950c8f339c8493d2298bb53e147c1bf8ce8cd2d54762253f90f43fb1c254"),
		expected:           decodeHexString(c, "d6bb1964e69c5612e58bff4660a5836704d7f14a3dd83bac427a464c8dcce60822c857f280c2540a5c4319b8f137f8cd5c9fb8bfa7f8ea75587695ada3b799bd"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_3(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "9aa2275145e252f9471fa1399eeaf84a7dac1590b6c12e7133843935587ee814"),
		nonce:              decodeHexString(c, "e50efcb1a4fac702f24df5047ef49d8c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "e05b0597bdde1998effb9702a20c792e8093c2896007f8777dc5933a6de49b10"),
		expected:           decodeHexString(c, "5ba6f7b65ec4c95d17cd029ad56a4fe29dd703c93313ac065974155964a7b9b0fe252bc2e865352e6a4caee090721a0eee0d6a7a0fd83c74feb728fdcbca4e94"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_4(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "f65ecbb21205f1486fd95f77a9acd61a392d9c9d80b8010c9989bb84ae31f064"),
		nonce:              decodeHexString(c, "32b04352bd345b8e46a5b77b308064b6"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "32d861ef5bccc90d393cc99b5c4550a41e2f0c2d234828235f06243d6126d15b"),
		expected:           decodeHexString(c, "524630ad63df6294b975d1fcf86b79506697c4b79668d382e7d83e30da06acbd97e16e256df73d680c5044e8343d6b88123c7c89482e93ef1a6c67f814cb998b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_5(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "3cade52468ee033f340cedf266f60e5dc4f446ce1c537509c3a25e776e2d054c"),
		nonce:              decodeHexString(c, "325e3c6bc90dab20178380bc97a92ea3"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "6c3a927d9f0620926f354f2b91298632bc526b0c99f215056f631e079726ad98"),
		expected:           decodeHexString(c, "1d2b132516354e9b0cbcf78812dda8fdd044af161c2ed5219df1a4e643dafabad1f2321ed09d680c278a2a6dadfb5a5c9cd3284c7e56262bb7077ef7751cc9bf"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_6(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "82bcbaf43005233f535ab04bdd9eb08f5524fb6999e9bb60c4b9501bb58faae5"),
		nonce:              decodeHexString(c, "b634f119617533242bc4e10cdc73c8cc"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "e0ea050554d4b7ab4faa51e384eb4a3dfcec08048a6eaa6d51e0fc956043ecb5"),
		expected:           decodeHexString(c, "17aef28a45c1ebeccfed991f526e560035d1c9e73de1217c2690e4e01b363c5148ccd80071143fc34df0eec73542d9937a226b13f16c2fcec968a41eb6a520d9"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_7(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "dc19df9e97759b8267a550eddb19c9ac936e881fe5f807d81bcc914b3c5f6389"),
		nonce:              decodeHexString(c, "1121d9752e5f882a707560bd0a449c59"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "1c615cd59622edf0e1a5bfdabd4c392e5dfde87de056ab833fe23854b26f3ae4"),
		expected:           decodeHexString(c, "62d466ee590a6c77279fd81c637d0c13b6ca886e7dd5380d5586428b40a636581752458adac6024cc63d5124b7f5400b3d254e4ddbbd48d2048789ca0e464f9c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_8(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "0062443385cdb8bcaed27ac3ef50a98d9346ff59f5e2242a2d0165d3a78aab58"),
		nonce:              decodeHexString(c, "448c174e316638eafa0bcc35fd5c599d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "ece3f65e9ee3875bac852cc68a8172bd02f3d70cf78607edd3c0686906dfbf7c"),
		expected:           decodeHexString(c, "d6e0cb062ca6f8ceaf3280d918062168df79c3a2a92817be76573f19b2d51515ee2070a78701ac41419c7af1f6d3b5cc3f7bc0f3c0cbe37c9c68258abf2b2b2c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_9(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "f9a0992037170cb0fca169742b0c7de2ee807b13701b29d4e49da04a00f204be"),
		nonce:              decodeHexString(c, "97869ed796b03cf4aef000ff750a17aa"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "3a221b7cda67d64afcdbb163f20db4584a39a0da8a70756fb249fd16fe960aee"),
		expected:           decodeHexString(c, "923b15e4f63d084c5993ca7202f0754da829a0ea426377197a4cbdfff49eff86515ebba839dc3fa7a72be79fdfe182c7c08aa83fc026f88e206c7b194d2f3ec5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_10(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "2c9d3934fd01418857c69aa5e650e66c2778c8e3c0d61801465285072a6c9628"),
		nonce:              decodeHexString(c, "c5b5065f97971e1f8d8f9070c6c18cbf"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "8550dea08a70965ddbac46d12f3445d1fbf3fcce233f540b23494b24b878ad03"),
		expected:           decodeHexString(c, "da8c3b43dcd7d393b69e4d023ef0ee57a49752cb16ef4faf8448ac674aac9cc3c438b98a8a6add54f509bd763e47d7a2eb5254009f6952d38e2bb6e05dc5a972"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_11(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "e6a5843f49e0a837bdf216e0644c8902effd8c6922d30eba8da3fd4537578e8c"),
		nonce:              decodeHexString(c, "d87b4ce9489aa9b6d1837d9c72ae9869"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "a90e89fd52d974d86c39e8d504a61a17eb08d3d89d97bc4f18de4ae28795757f"),
		expected:           decodeHexString(c, "9519154b87b2c9ac7be1e1ea12d5525d8e00e25b0528bfa0852e45e890197dcf3aba65e2812a42e3e925e2d8750ce59654c043cdd3a6c92d0914d030ce87a439"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_12(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "aad6e713433a4cce72958b1c69e22e67ebf0dee502abb5ce6b2a5ab35c0cef2c"),
		nonce:              decodeHexString(c, "6301b41a2c28b30cd357f08900d6ca75"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "a9dddaa088289241f65fa80ac6115979602798568956f1bb09340d78786ce3e8"),
		expected:           decodeHexString(c, "e8d49dbad6e089ad08441a105ad89016fb0361ad1b6dfb835b22836e5131e8a2c4bb2cee2a45e8181772194c29a82a89054df70d9701d277beafc8553c210258"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_13(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "749fad3edf120a72681e678b5d6836a73ed73612ad1b3757bb0054f761f211e9"),
		nonce:              decodeHexString(c, "cc84acb7c687636c1eb7701e6d45691c"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "468ece0720449193fdbee23dc5b24632770c44485b2bc19a2dc4e2ed3a45c935"),
		expected:           decodeHexString(c, "d50e347e74598ee3b279eb7f6be7834946e54496ce0c091af1501a30c6073ffdfd42c55f2094c816071af663daf38cf4515c4bb9a0d15a0f957709ad84086169"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed56_AES256_14(c *C) {
	s.testCTRAfterReseed56_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "ae8da96f3d66243a89ff454676d9bda5a39dadc723b119b48a03b7cd9da5b02e"),
		nonce:              decodeHexString(c, "bcd1d643680b50f1920513bbafd38b01"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "d296a0ab1df88a37542a5121a47409d5e20c48f48e3ac408d8a492f7cc21f282"),
		expected:           decodeHexString(c, "5fc1dbe67e396aed1d041eed15732a80cb3f50829bdc549c7959f26ed66ad407fa9398a58c7bf3257104e169ae6fb5dc18f0a185a8baa744dac4a114e16f111a"),
	})
}

func (s *drbgSuite) testCTRAfterReseed58_AES256(c *C, data *testData) {
	s.testCTRAfterReseed(c, 32, data)
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_0(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "7fc5c67c1e8eaebf19be6463c9ee13825b1c63bd38e58ce73a776887d95ff920"),
		nonce:              decodeHexString(c, "36b6aac81c45458d48e3a1a342ff667c"),
		personalization:    decodeHexString(c, "2196680672e2c4e164059cde6d2fe91ba3c396cf4b61b5e23fb1667816f9bda4"),
		entropyInputReseed: decodeHexString(c, "114475d8eeb771a0d9bad451245f3633e709592442e5005845d0ebafed5f680d"),
		expected:           decodeHexString(c, "cc7c9020a9b11501440464e3c306d38262c45838da3a0dd26552ee7a9edd9fc382d3f7b187e9fb370be97d9bf43466a551e9738929f38697c738bf267b664984"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_1(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "3af4df4e101056d22e9386a4f7d47a975a8e7b44e202e7a3d60a0c920c070f59"),
		nonce:              decodeHexString(c, "4fdbb787ede1f7041cd6c5a180c23726"),
		personalization:    decodeHexString(c, "f8519898a7173c7beee3406265243c0b06139c3cbcb47a6c4525c41f5cd079e9"),
		entropyInputReseed: decodeHexString(c, "8172999c005b5ea60ce12bfe0413d7c7974e55f1b8e0552139085e1ec9ae79fb"),
		expected:           decodeHexString(c, "fca17ab323f44a1f7bee2ac8400066eee2b02bfc434f63cc9fa3699b083b34ac7a9aa909b411c769cde12cab39b31d7077d41fa0dab0ab1abe8e7ee775511e3b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_2(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "e8ba22bc9d746b6a4ecf610bcaf197130cf62269dea684920bf1bbcf17660324"),
		nonce:              decodeHexString(c, "54afff3ab29557aaefbf4f2d7d34e94e"),
		personalization:    decodeHexString(c, "e021d4426537dd91590e354be4d96107a78db80ac4802fff384b529a3f8fa925"),
		entropyInputReseed: decodeHexString(c, "cceab6a26c170b689addc962be4c11a4fcfb472600e7a3e5c5e78f0ce8fa97f7"),
		expected:           decodeHexString(c, "d20454549422fbdc7708b047e2ecbd13bb4712e38ab2b0efc6800ce2d632acb2ac1436fc813d551134947d142d8421a91d1eb32150cbf99b266c552b215c20a7"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_3(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "cc0283b56b01af29df83617f12969e05bc95151bd6ea04337825891ac94798e9"),
		nonce:              decodeHexString(c, "825976f832796602d9afac19f9a45972"),
		personalization:    decodeHexString(c, "75aec9c32f40bda33902f1a21075775970f6a27844ae2a3429b5e186119ce917"),
		entropyInputReseed: decodeHexString(c, "7b273415d5bcacc9beba66599235b780a077f4a7ebba6aeddcdde583c20589cf"),
		expected:           decodeHexString(c, "e83757b19dc244f48dbf6aba22a8b24ade44dee959d017ffb4fe9771c2a6d28cc56e9449c9050f52b5a315ff7e45354352fc4b44621944dc7ca3a93fba7aa71c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_4(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "d4c9fa57d211f53dcd16b2f1812141ec3efe2d0bd425d5c1fd7e6d96a146db37"),
		nonce:              decodeHexString(c, "6473758b32848f04b86ccbcbd017f14b"),
		personalization:    decodeHexString(c, "a2698b2b6e58c23c3e82cc195e155164f4d8865392469a30874e549b0171a490"),
		entropyInputReseed: decodeHexString(c, "96b59a209fe54ce75a3f0d6f62f7e492aabc41584e1607463d161f99e98cbd88"),
		expected:           decodeHexString(c, "1b5bf3cfee33f7fd4b9a07f9bb98255b0bd47a3e8d6472af57602ab8b6abebd078df5aae7610533ae31738956c3e4ccd41104585655dab4cfcb32d37c81fb792"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_5(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "751be4104740cd52a1a515e43b80efda1738558de89e6f049ea194d7d8487f29"),
		nonce:              decodeHexString(c, "eec60a2978a490b2e6be71765e69c361"),
		personalization:    decodeHexString(c, "6c7a39822f61f4f17d0ae39099fdc820c635c69005bf04e4d13b18a188382140"),
		entropyInputReseed: decodeHexString(c, "3961b24f6427352d52f2dc45ee9d22814e7226567adcff950bf73d431ab8cb4a"),
		expected:           decodeHexString(c, "de87a4a765c5040d89743146696a6702d5cde905ebf2dd0f6540d53f5c8bd4fa1c3aa83b9c2b0edd72f857d59571ba508bd63d5f7ae30118e3e9688c606fd1cb"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_6(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "d401f6fc6daf7c003ccdfedeacd011e2c049963f662bfe185c631568a83c9eed"),
		nonce:              decodeHexString(c, "03735156236476104928df85c30774f0"),
		personalization:    decodeHexString(c, "175048786aa83e4dc8500b0118fcdace9174bb77ed8eaee4c55feba04534b09d"),
		entropyInputReseed: decodeHexString(c, "d54f61e2f153d3422f748706a4f407914b8478997518cd5f24f07b523bec5ce2"),
		expected:           decodeHexString(c, "a3a6910c6ffc121bdcfa6d29b7ee7872b537b3a3ab84a8c8a6d743b83de98dcf9be9dd506e51c5a5569e40eeffdb5789a05315aef595cf4401cdcd3116fe24fa"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_7(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "77230544aaa99a910369a2ef6798e2106246ab47296485c8f65a75303cf90b18"),
		nonce:              decodeHexString(c, "4937bc2c211ad7137797a0c4d0cd073a"),
		personalization:    decodeHexString(c, "1d2463a3266cac9840ed6b7c35f145654189e1e083222a4a281dab501e9923f9"),
		entropyInputReseed: decodeHexString(c, "2e1d69e466800fe9feab872b3e3d410459dc1791f2924241a00a585a6f94dce1"),
		expected:           decodeHexString(c, "2fed41010752d77f323c4bfbfa09b95296bcbe565b84b4a65d7eca938ed64e30f7e48e0c71b2ca0b6c08a0fe52d8a0cfd8558e58dc15e7d5610cc66c24225031"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_8(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "5009d818cc0384bfbc9c5dd81579c74f56e5490c56e5311152231e8a71720a13"),
		nonce:              decodeHexString(c, "1b164d66073467a8fcb1722c671b408e"),
		personalization:    decodeHexString(c, "1e56a2cdae5911d0d60baed7d49d3b3cb062c4cbfbbe31a56713beb37fb19fd4"),
		entropyInputReseed: decodeHexString(c, "2d754eda6de487a5b28284867557accd432c6386924b24d32263f607291737ec"),
		expected:           decodeHexString(c, "89fdece40ef81821bbb7beba79f1c1d68cb9bee9e9cf2f6c2b0bf42e0ca3a3c6659b2acbb5eee85ebacfdccf352022088b995ebc84cd24d3a19b832e4617cdad"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_9(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "d4dffe9c0d22f9e96cd2953c6104d79432848893b750eb2c2d738c157f2ad672"),
		nonce:              decodeHexString(c, "119f1204c0ab086229735c03cfaad1e7"),
		personalization:    decodeHexString(c, "811c7480afe7d9acfd5d46d85d01c2af4ec7802948a04d3f2d6c2be8ed80f2fb"),
		entropyInputReseed: decodeHexString(c, "59c7e94becdf5234b602903152eab24eed5f841759fd13a9b0da4ddf14821170"),
		expected:           decodeHexString(c, "8a2bdac0e3b9ecab36f495cdc92bfd0eb1978aa62604da0cf5884a0262123004958f07e70f3380de03a83868f431f8cefad7a14b5a3d9c4254e52f0ce991c0f1"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_10(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "0ea475261cad88913c57c98aaccc0369ba00e8b0676d1bfb5b30fc36b38a4a1d"),
		nonce:              decodeHexString(c, "3dd18e05b763fa1ca3881e7d92855c8e"),
		personalization:    decodeHexString(c, "1b95a4ba8393e52c9498279e8a0099013428c291d70cce70bea7f901a9b92828"),
		entropyInputReseed: decodeHexString(c, "e8ed162de3b51a68e2cd8df591b9c62259d24e31012b7938cd368c1a536ba9b3"),
		expected:           decodeHexString(c, "2c6cf1747bdac8f33351d6392daf4a2b32f5424b12f1fd8096b65b6a76398917c434ffedacdc6b2be1aa4ad6dcdad18932b638e3cb56deb72efd3b69cfd1b0bc"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_11(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "21893799016cfbe544d663800d6eaabc1d49fbe3e5bc4b958459cbe301ac678b"),
		nonce:              decodeHexString(c, "60059854b234f18be18c6e6e32c60d69"),
		personalization:    decodeHexString(c, "ab61e44f7aa31c1241a2a48c334d3fa95203a2f2102afc2909b627f83bd0eb0b"),
		entropyInputReseed: decodeHexString(c, "456a7d566bb488af9a8084e1961b610d05cd7ea57354b20f74c30818abdf011d"),
		expected:           decodeHexString(c, "6401c6afe7342ad95745269580fd3fb1a56c3c7d7f7d747d35e09501c0ef359e1dbbe1e0f5113421f889ad64ab6ee3bff21e38668f7ea16a800dd02179485bb4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_12(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "ad02f7d257da2a693bb3c544c066cdd77029fedf757f424ce212103e8349cd4f"),
		nonce:              decodeHexString(c, "6ee3dd909cbbcdc04623fd63e9154287"),
		personalization:    decodeHexString(c, "7f9b521d357ef486e36827fd0dc030bc5b485e7b44244555fbb924b20a88bc49"),
		entropyInputReseed: decodeHexString(c, "f5d103a7b483cffea2e09e5a5e849c436df08cb41e179e7d2c8537053ed5e71a"),
		expected:           decodeHexString(c, "6c212703a2783580f2fcb4dc48561352cf4bf23cf1ca3f42762923b9fad352d2ef63836760756721a3cd9153d5dc3e96a2d72acd5e6bd9d7a360f9ca05103c3f"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_13(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "a5cbff0d435897ecc9a787f805bbdb26b7ae8740f6b84a46513efb2cc97204ce"),
		nonce:              decodeHexString(c, "76a06ed1ede264a2ac16a58050e191f9"),
		personalization:    decodeHexString(c, "6f5cd73855fa56966d62d504e3211664edaaff51818e7a30904c3ae4b1a04a51"),
		entropyInputReseed: decodeHexString(c, "42e6b157ab34190f8260da9969a6fa35ec313d044fe5ef96f6ec497fe1e526d8"),
		expected:           decodeHexString(c, "6d5df4dec997c59a3630efdfa9d747ca8c82a86305612d40439162ad485d47c93ef44b884f69df3c4ab40a2e4ea63a455156415f31a31fddb6b18d9ae1f3cc1a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed58_AES256_14(c *C) {
	s.testCTRAfterReseed58_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "9252178c52d11248a056912f5100a3d7583ba41c91e531b755a93ac5e0c61bde"),
		nonce:              decodeHexString(c, "9ee982dfbd236075359a6f196bc0039f"),
		personalization:    decodeHexString(c, "58625e96e03c789f00659ec4cc7a13b31404b9916d8bb76f909f60c085fed9a4"),
		entropyInputReseed: decodeHexString(c, "deaa876e0c1bbdfd5dd7acf88b5320bf9bd1cc0a95c74616e6066cffd913dd7a"),
		expected:           decodeHexString(c, "340afc3154deffc2d8e2b9f9bb1d1c69576c6b355773e279f07e23267eea72d3cd7067f0c2dfa584f0fd1fe8e640058795bb24eab4e5da36148348703802b5fb"),
	})
}

func (s *drbgSuite) testCTRAfterReseed60_AES256(c *C, data *testData) {
	s.testCTRAfterReseed(c, 32, data)
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_0(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "d5559102cf8f234a89b6c48cbf473b1572a7d0c342d7b61adde3d6a0124d3991"),
		nonce:              decodeHexString(c, "5be948d054bb66e176b93fa848da0f51"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "8bd544ef239be98ff315261ad3a3e23a8400f1ebdcca65e0f46c7c661fc421a6"),
		expected:           decodeHexString(c, "e1bdd0bdb4d51b010b111e9088df562d216ca7371409d729f95250e8100f9753a60099a49408bb0065f99d59dce5081bd67cebd54c2b21fbf35184f26d1c4706"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_1(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "6b9dadcd05b1f2b4493355ec621bdbb0ebb67952337f3d372396319777477a70"),
		nonce:              decodeHexString(c, "34e62e1c2e741b4fd74b799c3f6fd9c1"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "24a9fc6393c8c3af6ba2ece51187d72980f40ad601f0395435c54edac642681f"),
		expected:           decodeHexString(c, "d2baa45967617b7d9a5056fa8b843d9f5c72b77ed951a1a4e43f2e88a63232bcf1cfb22718868a6d142af20d234a0b4a29f5f152d72ae60b9eb868953c0d46ad"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_2(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "55c465f279860ae0a30b374e5420b58f5c2fbb557928155bc049404c717d0148"),
		nonce:              decodeHexString(c, "d4137d0c64fd932057c99e9c488bc9e9"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "d0976462802628c6ed6320f6d88521228cc62eafd4a8e14984aacd0a30b21b1c"),
		expected:           decodeHexString(c, "c1406812252b57e793ce57132f0bf4b7e786a2b96ba284d76917288f0c79b5f52c591bef9b1231f982e142aae6e0cf63bff0e54a1c89345f591fe56d5a795f95"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_3(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "4071952b5c08ada347c7ad5eca7310963d0886c4f3076769c5ceb732985861c6"),
		nonce:              decodeHexString(c, "cc2dd3393509b4bb2542d2b69610d49e"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "bf9c1a5b5d9b7ce8f9e50c62daefef1904190552ae4abc222f8de865d3e3ee0d"),
		expected:           decodeHexString(c, "24fb483fb7c9ff58e2dc900d6334d3a3b62d26ea74e606b6dc7a9b1eb5079ffa0200d4f94795e1b2aeb58a481148f24832a8299216ea9c1724274ecfe2ed8d2f"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_4(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "8b1dbf309e22d7a792fa898b23db77c07338c5b5a90b89de5414b3d85bac8581"),
		nonce:              decodeHexString(c, "df1cc9e00dae202af131e81010443273"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "fa1fc8ff6aecf7ca00f3180e94fccbb055e3a2af28c27f66eaabb81351430b08"),
		expected:           decodeHexString(c, "5d34785040d4fdeb858ab1ca7c4bff23601fdfd91fe003e579e114a2e2a8f290e6c42b20c82322dca0f4c9abb634954d596d1d1bd1193734198352152e4eb817"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_5(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "a946beb38c95b63bb711f043b049ed94cb7d1e08018544a8fafd275313872a75"),
		nonce:              decodeHexString(c, "c858206dca843b65ad9e50a63ebc32bd"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "2d0fc4583542c4e9231482f66a522846bcdcb281d16eb07950a8a8595b200b9f"),
		expected:           decodeHexString(c, "2bbea16d110e8535fba89f4a9cec482c87d999982f6b05c15c4f4bcb740d1d43b90fb762aa8b506afa6d4c8b9676e3bceeb63db92245227c7366aa96970ce8ee"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_6(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "a7fda196a70f5bce96154f88d7a5137b17833a435f042a166d55504ec598b2a5"),
		nonce:              decodeHexString(c, "ff4e6e41b0720aa72c3493b3f207c258"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "ea04e60fc4ff309d0534b7ffb5b1a05499326baf5e0d57fcebd32be6ee508ccc"),
		expected:           decodeHexString(c, "b624e0594936ecc576aed106facf684012580b5cd7502c1625ad6e0323f64eff8b9176cebdd1f6ab7f399a4b71b8a910f912e12e7145bbb0bb47941066cc7ad5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_7(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "59701152798c85d20eda963c032b3d9efef8d7c714203ad44a3971e3a6efddd2"),
		nonce:              decodeHexString(c, "37963cf44dfe0387747e23fd2cd1256e"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "9c61d2946202c40e78370c46c3dbb4ece293099d88089788592cca1b4c49f79c"),
		expected:           decodeHexString(c, "14c142605f725bb594f200fbc709af892f0a324d41811fca6b81ec71c6a2ff1ee423de7e1421337760847e862670637546cf170735412fa262075219e102c240"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_8(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "9e4a3124f5c56e8369d51fa42bf66255130a3a30053427e0bb5d0366f18bdf47"),
		nonce:              decodeHexString(c, "55ddd182b956aaceaf92ed50c7ea7781"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "558f4ecf21687859935f9a25c2acdfc0099c693a86f1cefe62ef3b97334a3fdd"),
		expected:           decodeHexString(c, "02a7bfda634849ff49ceadb4ba679465cc457ff10735bba72b138c2127b5306b5af08fbbfa8fd417a67339bcaf93fdf417a26da6fe3295ddfafd0cb81a8eff3a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_9(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "762daf87b7a26de0bab5dba91101f898d1925a517a525339475dfb43a05a970d"),
		nonce:              decodeHexString(c, "f4c983088d46c475d49466ddf3356cd5"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "d392d2bd137acd801694fcedcfd7cb5cc8f5adf4b1cbf5e5a446c24e3692a260"),
		expected:           decodeHexString(c, "2b38ebe7a4b0ba7bc977d1e3852678f9a9ec78d99ee5c2e241dfdcd363dee1589fb66d8906eff7b492e2326931a6ea1159664978122ba6e208e49166f4811fa0"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_10(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "6ac8e3504024bd11aedea1fc286eb2ead9719837b1fb9568bca01c3ec9cc74d1"),
		nonce:              decodeHexString(c, "3b90c5c9dff9b052e4217b278aa64c3b"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "c4ccfacbf294b56e41f6b5d691ec36d9b69d70ed67829a8d2856da8593f2d068"),
		expected:           decodeHexString(c, "3b2860ce55e680fd964787e5d2bbf6fbdb7d5d8bc8dddf6629c588776004beb98314d1501d0e235beaff947627dd695d77b11b19e4feee86030e20479f3212df"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_11(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "a4f11dae4a6d515a7dbbc624dffa3dc16eefbf3be7207d9c1bfa2327e7889844"),
		nonce:              decodeHexString(c, "58a2e7a6f9cc543de0d7bbe82357d185"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "528c206ec59345ce4a8ffa8f5c85dcfa7847c1184a7984397869a1bd4ef6c146"),
		expected:           decodeHexString(c, "2097b57c802d2585ec192098eea7ad73ae11db7d284f75d2c31ff4b6cdbd3f42e2525a6518383d8d892a578af948425e1e60803afd8835c6f73b587aa78ac03b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_12(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "db544f7678e8e2f6c845e424414684d3cb6f2f67050df5af3ca5e41d5d834b42"),
		nonce:              decodeHexString(c, "c96c735705bec13ec9348e8f5db4555d"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "8374984b539263faf121459109e0f53a0303fa0f820b9bea4e35c5644a42bdd2"),
		expected:           decodeHexString(c, "d6b543a95c184b96583e9e5d477ed19e2dcdb84dba6921587fb14bc6dc1da5a7bf9f2aa4145d3d4f773f5393ec8b76c307f3f72ff1b5e0b03b322dda409a88b4"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_13(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "8681547745ec3c9bdce97528798ff35e259f0e649b525f492daef3a596fca9eb"),
		nonce:              decodeHexString(c, "927d4b92a9fd2d8a78a05f73c904a7c8"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "8ee9e576632b2f5a549e3dd0d99d0b267e07936e97223a49edf386bac7fdcdc4"),
		expected:           decodeHexString(c, "d1616f145a4efd1cdc6e688266d20da116ad9044ce4e4e42d7e6e78b65bb4ff15c8936df9f2ac65f43b5111750ece51412847ab3c0d63053545546fa66f57718"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed60_AES256_14(c *C) {
	s.testCTRAfterReseed60_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "734f99b0abc61ae231db7fb19ab2bc4bb903567d6b35dbaf76274c4a689ea2ae"),
		nonce:              decodeHexString(c, "3cbcda1dc5f77e82111125ac8c9d30b3"),
		personalization:    decodeHexString(c, ""),
		entropyInputReseed: decodeHexString(c, "a3756fff9903f323e573872104a1f4a2035dbd3a464ee5a5ee8db4f48ed0f1b7"),
		expected:           decodeHexString(c, "bfe7e205ac689d21ae2a2685e2eaf1ae0ad3e139881891c1d11b50594ddd22cf3d7ee4d5d4010b44c9b7a86f9b86d665a1b28ae21ac1d119ddb54e144ee8cb64"),
	})
}

func (s *drbgSuite) testCTRAfterReseed62_AES256(c *C, data *testData) {
	s.testCTRAfterReseed(c, 32, data)
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_0(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "fbae3ee02105a8a2353bbe9d806829cf78c8c312c782abf1554c6646cc37a1e5"),
		nonce:              decodeHexString(c, "b0479900a404e8e79c5f2fd7819232b9"),
		personalization:    decodeHexString(c, "54909fafc8f70428892f8d32ed51e95672892192d3955409e89c53dc6980d0af"),
		entropyInputReseed: decodeHexString(c, "aab36c9fab8bea6b9deb701fdf565d51e7a18b389808f8b938375d76f8657842"),
		expected:           decodeHexString(c, "8d1700f1f632df3400af0cc91c4d3d11da034993df5043cefa49fbc01784ed78099eec91d09395084df325ba02cdbd5b1abc64f9e347d81ae091ec081fe27d4c"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_1(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "7d4f1135a52bc86c13750fcc1e02d31d51af0573405e7ee1b61a5aec6f969ac9"),
		nonce:              decodeHexString(c, "c2b995988a6fdcbe043a415abb20f6d9"),
		personalization:    decodeHexString(c, "c81a7c88169f1ce64f5b8edd1eccfaa1ab853e487996c24d1368af364ffe8cb8"),
		entropyInputReseed: decodeHexString(c, "98772db6c038a6bfe328c9db0593bb12c71cb14d12ff5c5e6aa11201bd7e0658"),
		expected:           decodeHexString(c, "d5e5cf6a1d6728c50a958cfa9e3853a378f4b47d2a8bb841aef6bc55835143fe411860e4b3afbfc948ff87cf6e653336422dcc36b606560df66bcafd8302d7c5"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_2(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "03cd4e03108959a587a209765412c2deb88585369aa7280ad95abde3bc5e6b61"),
		nonce:              decodeHexString(c, "499c1512bc86f1b0eb1a0627dce2cc39"),
		personalization:    decodeHexString(c, "3356afd60365388538c277b87cc82f4d10a2fa6184ba36cac3f712d584d65dc2"),
		entropyInputReseed: decodeHexString(c, "61e05c8b87a35d5be47fed54ebf7543ddda13bcbb942d080718cceb07ed71808"),
		expected:           decodeHexString(c, "2efabe7f944ce49e27b86fda4e0dd9c46f119ca2541c8781dbec6be2cc74ce9ac208b24be5758375720f1c42e04187623d2ccdce7343d7c8c1244a66926e2866"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_3(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "e9a33feef5455be57a876a4eafd4febb02a313c77c64217ffb8c6fdb2c46fd9b"),
		nonce:              decodeHexString(c, "b0a2561d86f4127871dc6c0917fe01de"),
		personalization:    decodeHexString(c, "6231a999d00e07962d9826095ed0c249817d8647ae02d17c25057438eac5b506"),
		entropyInputReseed: decodeHexString(c, "7121a38b59f80a53641b0cebe2a6d1aeebf6e36696482d54a5b5bf0ad4490293"),
		expected:           decodeHexString(c, "165f56a01e61944eee87ce0c752a8a31117d6ead60c37beaa05d8a39ec6f42b6b9c90e471c840a6172facd9a1bd3db7d47709d665b49407a23020dafb897e853"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_4(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "05bdd4e143180e1be2d2a561b90559268e462ad56869f5f5d3480fc4bdd1e682"),
		nonce:              decodeHexString(c, "747d40d20f46a7f39ae52bab17ca61ce"),
		personalization:    decodeHexString(c, "403e35af4ffae9e3ee2d5f277e69b29d3f4a8dac36691ddb31507dda6fbe6650"),
		entropyInputReseed: decodeHexString(c, "5e4e32e94ed5e1dc894b7cf2857bf5e2218e46f2b69f8bf4555bcca61568af33"),
		expected:           decodeHexString(c, "a36846c720118736d0992a0afeb08530a2a3b68bed0c76076ea652509117943cee2f8f888f82c8c0405cffec84b2145821ca33686435afe145c04a49dfe1cd7a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_5(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "a635e43f3d97dc3511932aef96649862b46830df9ac0ebbc31b932ef51ada05b"),
		nonce:              decodeHexString(c, "b38d9926191b49eb995981bbfececba3"),
		personalization:    decodeHexString(c, "cd9838b07d041be31357fe9ebc01faf54731ccb90584d6c19523df3989cd866b"),
		entropyInputReseed: decodeHexString(c, "b8d111bacdf01b76b1482da9df8976ac34bfff06e10184065a339adab85a9ba5"),
		expected:           decodeHexString(c, "6e0fc3de87a9928477905b9b621f3f9f86fec7643007e4b560854fb2099c7caa58b862c7ff21980ed6f91f7867a6ec4870e2c32b34592809b8af1795807ca34b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_6(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "1ff0239922872d42ccb595f7bbdbe6ad86cb3952251e6e110b360fbdf419591e"),
		nonce:              decodeHexString(c, "a007416d48312457ffe08b438b54d929"),
		personalization:    decodeHexString(c, "70a3f7362745ba47a9d5d593817f096d881848f077e81f4339e79effa9919d82"),
		entropyInputReseed: decodeHexString(c, "56c968dc6d1703fb429355a0033b5b61e87a226167bb36017c70c4b177bb7ad8"),
		expected:           decodeHexString(c, "d037e0fa29bb5e48bb0e914c095a118cbdc46a81b8a5b68a84cf828ec39ca4909455cd002126ae1c3dad1279bf33fcc7d7475905d3b5f4b981c8fb158fe67c8d"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_7(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "e4c6e4010c7f5b79d067e035eae4d0f7e4c6b313c5d0f4dda3f07f770e637819"),
		nonce:              decodeHexString(c, "982ddfcbe5b655dbbf1f6e428c6ea216"),
		personalization:    decodeHexString(c, "d6117048bec591f90a4681d8e667707f9afcfe92f2eab86ef04b68a41de17f31"),
		entropyInputReseed: decodeHexString(c, "02939cd4004ae891b3bfcd210075c1a97876c0cef86d3ecfebe2c8caf80fe211"),
		expected:           decodeHexString(c, "a4182c34c1df827ec93ebfa0d515cc7d6f8ee22f3a769a30af0cd5ee7488ab68f70c4ce62ee314a047268c0045fb1b7d2584bdb646c3ed49c88510b4c54a676b"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_8(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "652db40eda989032ae1ed5ee901cbf950c31833a9d6f36e5159c2cc8245df3d6"),
		nonce:              decodeHexString(c, "728039b672c1149b9b48a118e67f738f"),
		personalization:    decodeHexString(c, "14a92992fcb0157780f8199af56ed1caec8ee624d9232da4a149c3d2a6e53494"),
		entropyInputReseed: decodeHexString(c, "8d6a04513dd5bdd3ee04dd9dc0d48edac041348bf69523a82b25860ec171add4"),
		expected:           decodeHexString(c, "6a7721ac74feae95b295883330e00de94280cd666c7391c7108667d1292dc88015f99130e561551b7241c9e5a06b476be944215b2366e664eb28e5c25b2fa984"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_9(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "983fb9e29d7cf7406e031a501a04f5efbae21d89d1ff7e3744c6251b1c6e51ec"),
		nonce:              decodeHexString(c, "d479f803128d7a87dc5f1880e9d182c5"),
		personalization:    decodeHexString(c, "0d58ac5ae040a369aa370c40deee131636097a1c7d2c262edf63bf939f342616"),
		entropyInputReseed: decodeHexString(c, "35aecbdd244a41972be4509a98ddc4d6467fa633e9353d9dd2c3442a30875039"),
		expected:           decodeHexString(c, "fc7b2cf9206a83b2a8d7edb178632a0c0c0bd3aa28b19a963fda7fab9d09928adedee6c37d3dd4b9f386529c6802d9a4f5f639dfa492bfad22d683b6c9fbbe6a"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_10(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "bd9e23eb4bb4dc2c3b58a7f4d32c8e932108fc7a2dc7a9f40dce671fc3fea1db"),
		nonce:              decodeHexString(c, "14732ec751a55666de4f16bac77d20ab"),
		personalization:    decodeHexString(c, "bf7e5c3fe8e3af805e61b2a2be73b237e95c5b93cf1e26d0435ab63414964740"),
		entropyInputReseed: decodeHexString(c, "62f6d8065fdb7279bf58a4008095f448519a21231c9b96d59272a9b5382b726e"),
		expected:           decodeHexString(c, "51f4374e6804ce989b4bf41e48de6bfd371f02343a07da6a7a651163f8a84d4ea7c705e0c5491dfe5eb8730dbe38d69d688b6d8351e9600c231cb7276d69dcee"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_11(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "ba06ca204754972a26ac9625c85c5c8094d8edb07f6f473ebb941b5771187a17"),
		nonce:              decodeHexString(c, "200980cc1668af5a4e54079619470be3"),
		personalization:    decodeHexString(c, "57d807d0a619f895ac683779e6c1f89baeebc93e17db5b5e80bddce5f85b002d"),
		entropyInputReseed: decodeHexString(c, "fa0e8f2a77c6c06a586809f3eae93aa7eac0a3d09c262a72a1886651ba25296e"),
		expected:           decodeHexString(c, "e022dbdfcd0188ed16413014f1707577c6af5b59ec41a46b983638c6a7e055b9fade91528c9e5c46d84a71d733a47cdde62f3fb47d3356029c4ec779fc885691"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_12(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "a7942a0d3b075461a29bb99343b10e1f10014f53097c3402744759d24baf439a"),
		nonce:              decodeHexString(c, "f268eb70dbdfa7ec611419eeb94bf884"),
		personalization:    decodeHexString(c, "f947754a3135bc1907f86f77f6f5224594b2c587193f7d86e343dbe8ae940af0"),
		entropyInputReseed: decodeHexString(c, "1f5725653a01fd3d3870a5874bb97e0910d48039589ceb80a0d41c2d3b07240c"),
		expected:           decodeHexString(c, "7248faac73e778281885473b0ad2ed56dc3c4ecb505a29c080c57dd507e56a50bfe9ce04c724ac7130cbfcf5227c8df51ad108fc5875ed13cfdd3eed7b95ed60"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_13(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "c17b592351ce97c2b9397d1d35f7849361ce0fbcc89d64ea24ee234489c87848"),
		nonce:              decodeHexString(c, "d02207b533ddfc79fd54e24747254268"),
		personalization:    decodeHexString(c, "6b8860de89dc493459c3e8221db10d601677eca93c86a43607c0ff558d26b704"),
		entropyInputReseed: decodeHexString(c, "267225f3a9aa0867a4be8e3e53015451cf58796ace50a36c657811e51bd52170"),
		expected:           decodeHexString(c, "2c075efcca1a603e609f35bdebf57556e87c1d418bbf2298788000b8254f70a44e98172e41c6ba51dc3521dc1969bc386c625ec0ef1289c42c3e27c52b4a2487"),
	})
}

func (s *drbgSuite) TestCTRAfterReseed62_AES256_14(c *C) {
	s.testCTRAfterReseed62_AES256(c, &testData{
		entropyInput:       decodeHexString(c, "f883b4bbea89cac2fd378559fe5790d7ad64dc6f5acc61ceecbc13bd971f6afb"),
		nonce:              decodeHexString(c, "3ab08948f01416317cebab29eb211d7b"),
		personalization:    decodeHexString(c, "d086057493500d75d93d9327b09c108ed9e62701794951c9b9fc77ef3872a555"),
		entropyInputReseed: decodeHexString(c, "2149693ad3bb60d8750e9f21ffc16b7178310afac1e2fa63334302cffa1c0a47"),
		expected:           decodeHexString(c, "e0598a33114cc183edb843415d697acadc91c39ba54100c7b14f79e67e47eb7f8d21cc1c5e4d744b329f717c88239035b91fd4b70e415f2697e9f9d436f3b001"),
	})
}
