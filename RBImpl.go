package main

var rbs = ConversionFields{
	newConversionImplF("Amount", 13, getAmount),
	newConversionImplF("Direction", 13, getDirection),

	newConversionImplF("PostingDate", 0, getDate),
	newConversionImpl("ItemNo", 18),
	newConversionImplF("PartnerAccNo", 5, getAccP2),
	newConversionImplF("PartnerAccBank", 5, getAccBank),
	newConversionImpl("PartnerAccName", 6),
	newConversionImplF("ValueDate", 1, getDate),
	newConversionImplF("ChargesAmount", 17, getAmount),
	newConversionImpl("ChargesCcy", 17),
	newConversionImpl("MovementTypeText", 7),
	newConversionImplF("BankCode", 2, getAccBank),
	newConversionImpl("BankCountryID", -1),
	newConversionImpl("BankName", -1),
	newConversionImplF("AccNoID", 2, getAccP2),
	newConversionImpl("AccNoCC", -1),
	newConversionImpl("AccName", 3),
	newConversionImpl("AccCcy", 14),
	newConversionImpl("AccCcyText", -1),
	newConversionImpl("Statistics1", 11),
	newConversionImpl("Statistics2", 10),
	newConversionImpl("Statistics3", 12),
	newConversionImpl("Statistics4", -1),
	newConversionImpl("Description1", 8),
}
