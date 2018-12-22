package main

var fakturoidVydane = ConversionFields{
	//	newConversionImpl("id",0),
	newConversionImpl("Kod", 3),
	newConversionImpl("VarSym", 4),
	newConversionImpl("NazFirmy", 13),
	newConversionImpl("Ulice", 14),
	newConversionImpl("Mesto", 16),
	newConversionImpl("Psc", 17),
	newConversionImpl("Ic", 19),
	newConversionImpl("Dic", 20),
	newConversionImplF("DatVyst", 27, getDateF),
	newConversionImplF("DuzpPuv", 28, getDateF),
	newConversionImplF("DatSplat", 30, getDateF),
	//	newConversionImpl("DatVyst",31),
	newConversionImplF("SumZklZakl", 43, getAmountF),
	newConversionImplF("SumZklCelkem", 43, getAmountF),

	newConversionImplF("SumDphZakl", 49, getAmountF),
	newConversionImplF("SumZklZaklP", 43, getAmountF),
	newConversionImplF("TypDokl", 43, getTypDokl),
}
