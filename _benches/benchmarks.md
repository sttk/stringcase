## 1.0.0

Date: 2026-08-08

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/stringcase
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkCamelCase-12                                       	 8297370	       142.8 ns/op	      80 B/op	       2 allocs/op
BenchmarkCobolCase-12                                       	 8074278	       148.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase-12                                       	 7926679	       151.2 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase-12                                       	 8066170	       150.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkPascalCase-12                                      	 8512885	       140.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkSnakeCase-12                                       	 7924692	       151.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase-12                                       	 7635226	       156.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsHead-12                    	 8451884	       141.4 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsTail-12                    	 8462102	       140.8 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsWord-12                    	 8520889	       140.5 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsPart-12                    	 8511387	       140.8 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsHead_withSeparators-12     	 7816942	       159.4 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsTail_withSeparators-12     	 7586220	       164.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsWord_withSeparators-12     	 7681776	       155.2 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsPart_withSeparators-12     	 7772131	       155.2 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsHead_withKeep-12           	 7475754	       158.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsTail_withKeep-12           	 7587164	       159.5 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsWord_withKeep-12           	 7626543	       165.2 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsPart_withKeep-12           	 7740266	       157.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsHead-12                    	 7143310	       182.2 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsTail-12                    	 8017194	       151.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsWord-12                    	 7509937	       160.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsPart-12                    	 7922592	       157.2 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsHead_withSeparators-12     	 7114582	       167.3 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsTail_withSeparators-12     	 7130990	       190.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsWord_withSeparators-12     	 6714057	       177.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsPart_withSeparators-12     	 7386193	       171.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsHead_withKeep-12           	 6857883	       168.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsTail_withKeep-12           	 7286779	       181.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsWord_withKeep-12           	 6792283	       175.3 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsPart_withKeep-12           	 7370794	       162.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsHead-12                    	 7476000	       156.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsTail-12                    	 7953217	       153.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsWord-12                    	 7505866	       158.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsPart-12                    	 7845207	       152.3 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsHead_withSeparators-12     	 7235162	       165.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsTail_withSeparators-12     	 7239914	       165.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsWord_withSeparators-12     	 7025714	       171.3 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsPart_withSeparators-12     	 7241395	       160.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsHead_withKeep-12           	 7043659	       169.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsTail_withKeep-12           	 7083126	       169.1 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsWord_withKeep-12           	 6724539	       174.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsPart_withKeep-12           	 7200871	       163.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsHead-12                    	 7847634	       155.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsTail-12                    	 8065178	       149.2 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsWord-12                    	 7804520	       155.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsPart-12                    	 7979148	       152.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsHead_withSeparators-12     	 7109287	       164.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsTail_withSeparators-12     	 7165911	       163.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsWord_withSeparators-12     	 7102852	       167.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsPart_withSeparators-12     	 7373737	       162.2 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsHead_withKeep-12           	 7352380	       163.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsTail_withKeep-12           	 7357119	       163.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsWord_withKeep-12           	 7101054	       171.2 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsPart_withKeep-12           	 7355962	       167.3 ns/op	     120 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsHead-12                   	 8370392	       142.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsTail-12                   	 8370316	       141.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsWord-12                   	 8599782	       138.8 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsPart-12                   	 8661070	       138.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsHead_withSeparators-12    	 7900195	       152.1 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsTail_withSeparators-12    	 7752945	       154.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsWord_withSeparators-12    	 7910958	       155.5 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsPart_withSeparators-12    	 7625026	       154.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsHead_withKeep-12          	 7895022	       152.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsTail_withKeep-12          	 7781599	       157.3 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsWord_withKeep-12          	 7577026	       155.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsPart_withKeep-12          	 7889040	       152.8 ns/op	      80 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsHead-12                    	 7645876	       156.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsTail-12                    	 7992746	       153.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsWord-12                    	 7514497	       158.3 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsPart-12                    	 7821312	       155.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsHead_withSeparators-12     	 7001104	       168.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsTail_withSeparators-12     	 7251379	       166.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsWord_withSeparators-12     	 6894979	       171.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsPart_withSeparators-12     	 7425940	       163.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsHead_withKeep-12           	 7116295	       167.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsTail_withKeep-12           	 7217629	       166.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsWord_withKeep-12           	 6983959	       175.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsPart_withKeep-12           	 7361854	       164.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsHead-12                    	 7366144	       161.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsTail-12                    	 7671272	       158.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsWord-12                    	 7464363	       162.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsPart-12                    	 7669647	       157.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsHead_withSeparators-12     	 6935749	       170.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsTail_withSeparators-12     	 7053303	       169.1 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsWord_withSeparators-12     	 6844314	       176.1 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsPart_withSeparators-12     	 7183777	       165.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsHead_withKeep-12           	 6830284	       173.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsTail_withKeep-12           	 6873212	       174.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsWord_withKeep-12           	 6621453	       177.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsPart_withKeep-12           	 7113187	       167.2 ns/op	     120 B/op	       2 allocs/op
PASS
ok  	github.com/sttk/stringcase	125.983s
```

## 0.3.0

Date: 2026-08-08

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/stringcase
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkCamelCase-12                                       	 8238853	       143.1 ns/op	      80 B/op	       2 allocs/op
BenchmarkCobolCase-12                                       	 7467542	       159.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase-12                                       	 7424259	       160.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase-12                                       	 7874016	       151.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkPascalCase-12                                      	 8190222	       145.5 ns/op	      80 B/op	       2 allocs/op
BenchmarkSnakeCase-12                                       	 7219948	       165.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase-12                                       	 7133396	       167.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsHead-12                    	 8354859	       142.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsTail-12                    	 8381600	       142.4 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsWord-12                    	 8332638	       142.4 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsPart-12                    	 8369208	       142.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsHead_withSeparators-12     	 7717772	       155.2 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsTail_withSeparators-12     	 7702347	       154.2 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsWord_withSeparators-12     	 7689754	       154.0 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsPart_withSeparators-12     	 7677118	       154.8 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsHead_withKeep-12           	 7505181	       157.2 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsTail_withKeep-12           	 7511365	       159.1 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsWord_withKeep-12           	 7592278	       156.0 ns/op	      80 B/op	       2 allocs/op
BenchmarkCamelCase_nonAlphabetsAsPart_withKeep-12           	 7650925	       157.3 ns/op	      80 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsHead-12                    	 7283523	       164.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsTail-12                    	 7444514	       160.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsWord-12                    	 7194865	       164.3 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsPart-12                    	 7166988	       162.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsHead_withSeparators-12     	 6928543	       171.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsTail_withSeparators-12     	 6764898	       175.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsWord_withSeparators-12     	 6578619	       178.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsPart_withSeparators-12     	 7041691	       169.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsHead_withKeep-12           	 6888884	       171.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsTail_withKeep-12           	 6733850	       176.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsWord_withKeep-12           	 6568836	       179.2 ns/op	     120 B/op	       2 allocs/op
BenchmarkCobolCase_nonAlphabetsAsPart_withKeep-12           	 7031632	       168.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsHead-12                    	 7320552	       161.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsTail-12                    	 7501370	       159.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsWord-12                    	 7151478	       164.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsPart-12                    	 7511533	       159.2 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsHead_withSeparators-12     	 6976078	       170.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsTail_withSeparators-12     	 6761168	       175.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsWord_withSeparators-12     	 6652554	       178.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsPart_withSeparators-12     	 7083810	       168.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsHead_withKeep-12           	 6856740	       174.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsTail_withKeep-12           	 6693003	       180.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsWord_withKeep-12           	 6328858	       181.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkKebabCase_nonAlphabetsAsPart_withKeep-12           	 7045820	       169.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsHead-12                    	 7601318	       156.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsTail-12                    	 7863204	       150.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsWord-12                    	 7596594	       157.1 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsPart-12                    	 7830248	       150.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsHead_withSeparators-12     	 7179931	       165.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsTail_withSeparators-12     	 7110862	       166.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsWord_withSeparators-12     	 6895914	       171.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsPart_withSeparators-12     	 7268700	       163.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsHead_withKeep-12           	 7011969	       169.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsTail_withKeep-12           	 7062927	       168.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsWord_withKeep-12           	 6849974	       173.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkMacroCase_nonAlphabetsAsPart_withKeep-12           	 7262916	       164.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsHead-12                   	 8199854	       144.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsTail-12                   	 8176182	       145.8 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsWord-12                   	 8120142	       144.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsPart-12                   	 8196525	       144.9 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsHead_withSeparators-12    	 7631168	       156.4 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsTail_withSeparators-12    	 7344124	       157.8 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsWord_withSeparators-12    	 7495515	       157.1 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsPart_withSeparators-12    	 7573231	       156.2 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsHead_withKeep-12          	 7529244	       160.2 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsTail_withKeep-12          	 7452739	       159.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsWord_withKeep-12          	 7419571	       158.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkPascalCase_nonAlphabetsAsPart_withKeep-12          	 7476789	       158.2 ns/op	      80 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsHead-12                    	 7069972	       167.3 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsTail-12                    	 7194195	       165.2 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsWord-12                    	 7012906	       169.1 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsPart-12                    	 7225486	       164.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsHead_withSeparators-12     	 6842712	       174.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsTail_withSeparators-12     	 6755245	       176.7 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsWord_withSeparators-12     	 6479038	       182.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsPart_withSeparators-12     	 6930892	       171.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsHead_withKeep-12           	 6664258	       177.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsTail_withKeep-12           	 6668872	       179.1 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsWord_withKeep-12           	 6478582	       184.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkSnakeCase_nonAlphabetsAsPart_withKeep-12           	 6837478	       173.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsHead-12                    	 6978230	       171.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsTail-12                    	 7027615	       168.0 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsWord-12                    	 6931922	       171.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsPart-12                    	 7130389	       167.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsHead_withSeparators-12     	 6465681	       182.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsTail_withSeparators-12     	 6412150	       185.5 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsWord_withSeparators-12     	 5826064	       188.3 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsPart_withSeparators-12     	 6784014	       176.1 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsHead_withKeep-12           	 6561984	       181.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsTail_withKeep-12           	 6420482	       184.6 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsWord_withKeep-12           	 6220905	       190.9 ns/op	     120 B/op	       2 allocs/op
BenchmarkTrainCase_nonAlphabetsAsPart_withKeep-12           	 6719670	       178.1 ns/op	     120 B/op	       2 allocs/op
PASS
ok  	github.com/sttk/stringcase	123.916s
```
