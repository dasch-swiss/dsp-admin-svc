# Changelog

## [1.5.0](https://www.github.com/dasch-swiss/dsp-admin-svc/compare/v1.4.1...v1.5.0) (2021-08-10)


### Enhancements

* **admin-service & ui:** basic role based access control (DSP-1843) ([#6](https://www.github.com/dasch-swiss/dsp-admin-svc/issues/6)) ([8f13f65](https://www.github.com/dasch-swiss/dsp-admin-svc/commit/8f13f6560a83cf0f0bca09682e96aaa3abc1e1b4))
* **projects-api:** permissions (DSP-1787) ([#2](https://www.github.com/dasch-swiss/dsp-admin-svc/issues/2)) ([c54ec74](https://www.github.com/dasch-swiss/dsp-admin-svc/commit/c54ec74189ebd0d9ba5a6b352c944b75e0f4d9f9))

### [1.4.1](https://www.github.com/dasch-swiss/dasch-service-platform/compare/v1.4.0...v1.4.1) (2021-07-14)


### Bug Fixes

* metadata page gives 404 error after release 1.4.0 (DSP-1793) ([#102](https://www.github.com/dasch-swiss/dasch-service-platform/issues/102)) ([4090965](https://www.github.com/dasch-swiss/dasch-service-platform/commit/40909657a88a5213495e6c72c570f3120e03db13))


### Maintenance

* **admin-service:** removes organization and user implementations ([#100](https://www.github.com/dasch-swiss/dasch-service-platform/issues/100)) ([26107c4](https://www.github.com/dasch-swiss/dasch-service-platform/commit/26107c4b8aa604017f96eb33901c24916f3fafcc))

## [1.4.0](https://www.github.com/dasch-swiss/dasch-service-platform/compare/v1.3.0...v1.4.0) (2021-07-13)


### Bug Fixes

* **metadata:** remove EULA link ([#99](https://www.github.com/dasch-swiss/dasch-service-platform/issues/99)) ([1ff6170](https://www.github.com/dasch-swiss/dasch-service-platform/commit/1ff617038e43be9b1990fec4f8a1f2d775974f83))


### Enhancements

* **admin-gui:** admin service project gui (DSP-1738) ([#91](https://www.github.com/dasch-swiss/dasch-service-platform/issues/91)) ([8e90261](https://www.github.com/dasch-swiss/dasch-service-platform/commit/8e902611e35472fb1423e40ae41f5212f5d92ce8))
* **admin-gui:** keycloak authentication (DSP-1760) ([#95](https://www.github.com/dasch-swiss/dasch-service-platform/issues/95)) ([7e6b7b5](https://www.github.com/dasch-swiss/dasch-service-platform/commit/7e6b7b59c67e5ebb7628f07182e05f667c649e90))
* **admin-service:** admin service authorization (DSP-1772) ([#96](https://www.github.com/dasch-swiss/dasch-service-platform/issues/96)) ([13cd89a](https://www.github.com/dasch-swiss/dasch-service-platform/commit/13cd89aafba4eb33aa9380bfb7306275a0431698))
* write generic Go SPA and API server (DSP-1759) ([#93](https://www.github.com/dasch-swiss/dasch-service-platform/issues/93)) ([4edb145](https://www.github.com/dasch-swiss/dasch-service-platform/commit/4edb14535d678cb469f84c47ee5537b377f58bbc))


### Documentation

* updates slate docs and README ([#97](https://www.github.com/dasch-swiss/dasch-service-platform/issues/97)) ([6816c3a](https://www.github.com/dasch-swiss/dasch-service-platform/commit/6816c3a7ce8211418a54527b643d87fced0a89e5))

## [1.3.0](https://www.github.com/dasch-swiss/dasch-service-platform/compare/v1.2.0...v1.3.0) (2021-06-15)


### Bug Fixes

* **metadata-service:** fix cookie domain ([#80](https://www.github.com/dasch-swiss/dasch-service-platform/issues/80)) ([5a09ee7](https://www.github.com/dasch-swiss/dasch-service-platform/commit/5a09ee783bdab95fac8a584acfda40565f228a78))
* **metadata:** improve header and footer display on desktop ([#88](https://www.github.com/dasch-swiss/dasch-service-platform/issues/88)) ([3d9ec6e](https://www.github.com/dasch-swiss/dasch-service-platform/commit/3d9ec6eea617aed1a1072456aa3bc53059d4175f))
* **metadata:** relocate go-back-button for mobile ([#86](https://www.github.com/dasch-swiss/dasch-service-platform/issues/86)) ([58306f3](https://www.github.com/dasch-swiss/dasch-service-platform/commit/58306f3be2643cdcd882c8ea43ceefe7975775e7))


### Documentation

* **project:** updates slate docs for the project service ([#85](https://www.github.com/dasch-swiss/dasch-service-platform/issues/85)) ([42834c7](https://www.github.com/dasch-swiss/dasch-service-platform/commit/42834c7d9970dd7ddf436432cd98e12dec2c35fa))


### Maintenance

* **build:** add version.txt filegroup ([#84](https://www.github.com/dasch-swiss/dasch-service-platform/issues/84)) ([c119a22](https://www.github.com/dasch-swiss/dasch-service-platform/commit/c119a223a04e6c9acbc2ac5b5763310f809a71bb))
* **metadata:** add Fathom analytics ([#89](https://www.github.com/dasch-swiss/dasch-service-platform/issues/89)) ([d72d402](https://www.github.com/dasch-swiss/dasch-service-platform/commit/d72d402f01a839e49fa673baf4f15343abc20fb0))
* **metadata:** remove Matomo tracker ([#83](https://www.github.com/dasch-swiss/dasch-service-platform/issues/83)) ([3d04ea3](https://www.github.com/dasch-swiss/dasch-service-platform/commit/3d04ea35084ea2ff773e8a277374cb9d383925af))
* **project:** clean up error handling ([#81](https://www.github.com/dasch-swiss/dasch-service-platform/issues/81)) ([a73229c](https://www.github.com/dasch-swiss/dasch-service-platform/commit/a73229ccf8819e5254db098dd945137bc9be8c66))


### Enhancements

* **metadata-service:** add consent form (DSP-1704) ([#75](https://www.github.com/dasch-swiss/dasch-service-platform/issues/75)) ([cf05f58](https://www.github.com/dasch-swiss/dasch-service-platform/commit/cf05f585e90112c2773638c85f4dd8849d79ff17))
* **metadata-service:** add version number display ([#78](https://www.github.com/dasch-swiss/dasch-service-platform/issues/78)) ([f34e922](https://www.github.com/dasch-swiss/dasch-service-platform/commit/f34e9223091e5bd51d1d3b4af93d5a14a10b7d57))
* **metadata-service:** set domain cookie ([#79](https://www.github.com/dasch-swiss/dasch-service-platform/issues/79)) ([dee3f56](https://www.github.com/dasch-swiss/dasch-service-platform/commit/dee3f5640e0eafcc6105e0982c3701b499c65478))
* **metadata:** cookies banner restyled ([#82](https://www.github.com/dasch-swiss/dasch-service-platform/issues/82)) ([2ba6315](https://www.github.com/dasch-swiss/dasch-service-platform/commit/2ba63158eda751ad2948dd6b148560ca62f776bf))
* **metadata:** remove GA + update consent form (DSP-1742) ([#90](https://www.github.com/dasch-swiss/dasch-service-platform/issues/90)) ([3a13122](https://www.github.com/dasch-swiss/dasch-service-platform/commit/3a131223c92e0dc7e09c1394bd0582d02380878e))

## [1.2.0](https://www.github.com/dasch-swiss/dasch-service-platform/compare/v1.1.0...v1.2.0) (2021-06-02)


### Enhancements

* **valueobject:** adds method to identifier value object to create a valid uuid from a provided byte array ([#69](https://www.github.com/dasch-swiss/dasch-service-platform/issues/69)) ([6cfe24d](https://www.github.com/dasch-swiss/dasch-service-platform/commit/6cfe24d36fdf4ec790f1e280837fc3a8a702ac4e))


### Maintenance

* **metadata-service:** add Google Analytics ([#71](https://www.github.com/dasch-swiss/dasch-service-platform/issues/71)) ([4b6916c](https://www.github.com/dasch-swiss/dasch-service-platform/commit/4b6916ce3e0b9ee5ebfb27fde3f9291ca47f8f5c))
* **metadata-service:** enable GA on PROD only ([#72](https://www.github.com/dasch-swiss/dasch-service-platform/issues/72)) ([e341dd5](https://www.github.com/dasch-swiss/dasch-service-platform/commit/e341dd539815f2258c417b8ad92055aa838691a0))
* **metadata-service:** enable GA on PROD only ([#72](https://www.github.com/dasch-swiss/dasch-service-platform/issues/72)) ([42fb5aa](https://www.github.com/dasch-swiss/dasch-service-platform/commit/42fb5aa942e11094fbb1fe4479d1a9c998495ad4))

## [1.1.0](https://www.github.com/dasch-swiss/dasch-service-platform/compare/v1.0.0...v1.1.0) (2021-06-01)


### Bug Fixes

* **metadata-service:** fix Matomo config ([#66](https://www.github.com/dasch-swiss/dasch-service-platform/issues/66)) ([733d882](https://www.github.com/dasch-swiss/dasch-service-platform/commit/733d882b2cbe8953254050210f9cad639f2f3335))
* **metadata-service:** fix Matomo configuration (DSP-1671) ([#65](https://www.github.com/dasch-swiss/dasch-service-platform/issues/65)) ([d95ffec](https://www.github.com/dasch-swiss/dasch-service-platform/commit/d95ffec98e34d270497e1412930f5ab588465a36))
* **metadata-service:** snackbar handling + responsiveness ([#55](https://www.github.com/dasch-swiss/dasch-service-platform/issues/55)) ([77762ee](https://www.github.com/dasch-swiss/dasch-service-platform/commit/77762eef47588950b2e29191b08ac40a2b2b831c))


### Maintenance

* **metadata-service:** enable Matomo on PROD only ([#68](https://www.github.com/dasch-swiss/dasch-service-platform/issues/68)) ([64b1de3](https://www.github.com/dasch-swiss/dasch-service-platform/commit/64b1de3967f0e4b3b85b06f6eb76ede4f57b5bb9))
* **project:** add more descriptive method comments (DSP-1680) ([#61](https://www.github.com/dasch-swiss/dasch-service-platform/issues/61)) ([45ac2d1](https://www.github.com/dasch-swiss/dasch-service-platform/commit/45ac2d1177b347428e9eeb6a574e020da06e3647))


### Enhancements

* **metadata-service:** add Mamoto tracker (DSP-1671) ([#62](https://www.github.com/dasch-swiss/dasch-service-platform/issues/62)) ([4096b92](https://www.github.com/dasch-swiss/dasch-service-platform/commit/4096b928176fc7cbdac1d306daa925dde4ddaa2d))
* **metadata-service:** add sitemap generation ([#67](https://www.github.com/dasch-swiss/dasch-service-platform/issues/67)) ([3ee4449](https://www.github.com/dasch-swiss/dasch-service-platform/commit/3ee4449cf86f771081556468f238b37c63f12e2a))
* **project:** add migration endpoint (DSP-1666) ([#58](https://www.github.com/dasch-swiss/dasch-service-platform/issues/58)) ([58a4544](https://www.github.com/dasch-swiss/dasch-service-platform/commit/58a4544b168b814ec47078daadcf51a1a8156477))
* **project:** unique and valid short codes (DSP-1685) ([#63](https://www.github.com/dasch-swiss/dasch-service-platform/issues/63)) ([6742a6f](https://www.github.com/dasch-swiss/dasch-service-platform/commit/6742a6f6795e4b29f06a77d691133732485d2398))

## [1.0.0](https://www.github.com/dasch-swiss/dasch-service-platform/compare/v0.0.3...v1.0.0) (2021-05-25)


### ⚠ BREAKING CHANGES

* release please action (#53)

### Bug Fixes

* **metadata-service:** attributions & spatial coverage display (DSP-1598) ([#46](https://www.github.com/dasch-swiss/dasch-service-platform/issues/46)) ([460cd11](https://www.github.com/dasch-swiss/dasch-service-platform/commit/460cd11d35dee9c0059411679c2ce22eed241ea9))
* **metadata-service:** distinguish API call in envs ([#35](https://www.github.com/dasch-swiss/dasch-service-platform/issues/35)) ([4df0392](https://www.github.com/dasch-swiss/dasch-service-platform/commit/4df039252a9dcc521704a82af4c2d6a8830c52d3))
* **metadata-service:** distinguish API call in envs ([#36](https://www.github.com/dasch-swiss/dasch-service-platform/issues/36)) ([65eed67](https://www.github.com/dasch-swiss/dasch-service-platform/commit/65eed6700310cf4ba3a3d21be723c66f589bcbeb))
* **metadata-service:** minor bugs and data fix + look and feel improvement (DSP-1598) ([#40](https://www.github.com/dasch-swiss/dasch-service-platform/issues/40)) ([bf6fa48](https://www.github.com/dasch-swiss/dasch-service-platform/commit/bf6fa4860f0c4b1a5fef01e45aeb512f672475a1))
* release please action ([#53](https://www.github.com/dasch-swiss/dasch-service-platform/issues/53)) ([c56f63b](https://www.github.com/dasch-swiss/dasch-service-platform/commit/c56f63b07df147aa9b8ccf6859c21d6109927e48))
* Svelte + Rollup configuration (BS-14) ([#4](https://www.github.com/dasch-swiss/dasch-service-platform/issues/4)) ([8cab324](https://www.github.com/dasch-swiss/dasch-service-platform/commit/8cab324614a77ff6f7acf6b9fa0f6356670365f8))


### Maintenance

* **make:** add target to run meta-service inside docker ([#29](https://www.github.com/dasch-swiss/dasch-service-platform/issues/29)) ([eff2038](https://www.github.com/dasch-swiss/dasch-service-platform/commit/eff20381ad2ddab4cab11cbb30eb15fb68689af3))
* **metadata-service:** add more projects to db (BS-39) ([#12](https://www.github.com/dasch-swiss/dasch-service-platform/issues/12)) ([6d0998a](https://www.github.com/dasch-swiss/dasch-service-platform/commit/6d0998aa6f842918b139016ec1bb58976351161a))
* **metadata-service:** change default port to 3000 ([#34](https://www.github.com/dasch-swiss/dasch-service-platform/issues/34)) ([be05d38](https://www.github.com/dasch-swiss/dasch-service-platform/commit/be05d3829d16cd89d231d4c688caf00ed39c107a))
* **metadata-service:** update database ([#16](https://www.github.com/dasch-swiss/dasch-service-platform/issues/16)) ([58df498](https://www.github.com/dasch-swiss/dasch-service-platform/commit/58df49852b1da2fb16d096ade68b297db031b21e))
* **metadata:** add BMF, fix links, fix typos ([#37](https://www.github.com/dasch-swiss/dasch-service-platform/issues/37)) ([643f0d4](https://www.github.com/dasch-swiss/dasch-service-platform/commit/643f0d4125b73db73b7b176d740a228359853ae5))
* rename folders to match service and app boundary ([9f6af85](https://www.github.com/dasch-swiss/dasch-service-platform/commit/9f6af85351d1bb7909617a870458993f44b163d8))
* **services:** monorepo scaffolding (BS-6) ([#2](https://www.github.com/dasch-swiss/dasch-service-platform/issues/2)) ([a644a75](https://www.github.com/dasch-swiss/dasch-service-platform/commit/a644a75ff489cc95fe9369a2043f5e8b013b1ae9))
* **shared:** move shared stuff ([#9](https://www.github.com/dasch-swiss/dasch-service-platform/issues/9))(BS-29) ([6125424](https://www.github.com/dasch-swiss/dasch-service-platform/commit/6125424e30050503cd2484a7e714225529540d62))


### Documentation

* **metadata-service:** refine data model (BS-41) ([#18](https://www.github.com/dasch-swiss/dasch-service-platform/issues/18)) ([77ff52f](https://www.github.com/dasch-swiss/dasch-service-platform/commit/77ff52f4d6c30032a51763339187e412409d4887))
* **metadata-service:** update API specification to improve multi language text representation (DSP-1419) ([#42](https://www.github.com/dasch-swiss/dasch-service-platform/issues/42)) ([72d476c](https://www.github.com/dasch-swiss/dasch-service-platform/commit/72d476ca0b7242967f496fc363829be47972a684))


### Enhancements

* **admin-user-service:** add signup (DSP-1449) ([#26](https://www.github.com/dasch-swiss/dasch-service-platform/issues/26)) ([39726cf](https://www.github.com/dasch-swiss/dasch-service-platform/commit/39726cfc54c3fd1ebbbdfbafc3b82be6f6d18bed))
* **docs:** add slate docs ([#25](https://www.github.com/dasch-swiss/dasch-service-platform/issues/25)) ([2d2cfa1](https://www.github.com/dasch-swiss/dasch-service-platform/commit/2d2cfa10a3c17f8e33898ed09b58b708b8a0135e))
* **metadata-frontend:** change UI styling to mobile-first ([#10](https://www.github.com/dasch-swiss/dasch-service-platform/issues/10)) ([61826a4](https://www.github.com/dasch-swiss/dasch-service-platform/commit/61826a4b03e23a0982bde08c734b7a25fb291366))
* **metadata-module:** add routing (BS-20) ([#19](https://www.github.com/dasch-swiss/dasch-service-platform/issues/19)) ([4941680](https://www.github.com/dasch-swiss/dasch-service-platform/commit/49416802ea289e8ef5acee91d2ab11c5c8312353))
* **metadata-service:** add basic search functionality (DSP-1482) ([#22](https://www.github.com/dasch-swiss/dasch-service-platform/issues/22)) ([1632e50](https://www.github.com/dasch-swiss/dasch-service-platform/commit/1632e50f6e00d8148a2298683b3b1c0748888f40))
* **metadata-service:** add header and footer components (BS-15) ([#6](https://www.github.com/dasch-swiss/dasch-service-platform/issues/6)) ([7198b65](https://www.github.com/dasch-swiss/dasch-service-platform/commit/7198b653920554e8ddcd70f62ea4c2ad7a919e91))
* **metadata-service:** add json-server with example metadata (BS-11) ([#3](https://www.github.com/dasch-swiss/dasch-service-platform/issues/3)) ([56d489c](https://www.github.com/dasch-swiss/dasch-service-platform/commit/56d489cfb5c22723431ee2f9392add69371a0bba))
* **metadata-service:** add pagination component (BS-40) ([#14](https://www.github.com/dasch-swiss/dasch-service-platform/issues/14)) ([a8c9b15](https://www.github.com/dasch-swiss/dasch-service-platform/commit/a8c9b15feb6bbcc6621cb243b9ab75887428c24f))
* **metadata-service:** add projects (DSP-1454) ([#30](https://www.github.com/dasch-swiss/dasch-service-platform/issues/30)) ([fef0b90](https://www.github.com/dasch-swiss/dasch-service-platform/commit/fef0b90b7b017c329d03cb5b64a1345d52531a75))
* **metadata-service:** add projects list content page (BS-9) ([#5](https://www.github.com/dasch-swiss/dasch-service-platform/issues/5)) ([0aa02c5](https://www.github.com/dasch-swiss/dasch-service-platform/commit/0aa02c5fbe4e393f9d0c897073d8116dafa747ba))
* **metadata-service:** adjust Project Page (DSP-1550) ([#27](https://www.github.com/dasch-swiss/dasch-service-platform/issues/27)) ([32b121a](https://www.github.com/dasch-swiss/dasch-service-platform/commit/32b121aecfd5fc3cb9ac78cb18ae0ea575223a53))
* **metadata-service:** FE improvements (DSP-1578) ([#33](https://www.github.com/dasch-swiss/dasch-service-platform/issues/33)) ([cd84187](https://www.github.com/dasch-swiss/dasch-service-platform/commit/cd841870371f2d61fd2222a4405dd8a8b6f23a76))
* **metadata-service:** implement fake metadata backend in go (DSP-1579) ([#32](https://www.github.com/dasch-swiss/dasch-service-platform/issues/32)) ([c0e4a32](https://www.github.com/dasch-swiss/dasch-service-platform/commit/c0e4a320e1f0e9311cf7def031917b14f7ca2dce))
* **metadata-service:** implement Project Page (DSP-1488) ([#23](https://www.github.com/dasch-swiss/dasch-service-platform/issues/23)) ([752df8f](https://www.github.com/dasch-swiss/dasch-service-platform/commit/752df8ff0973fea281ff95214bd9e78b3564f4d4))
* **metadata-service:** prepare for production release (DSP-1578) ([#38](https://www.github.com/dasch-swiss/dasch-service-platform/issues/38)) ([e351f45](https://www.github.com/dasch-swiss/dasch-service-platform/commit/e351f45ed3a61ba5153c8eed888cb1380d6edff4))
* **metadata-service:** replace svelte-spa-router with svelte-routing (DSP-1638) ([#47](https://www.github.com/dasch-swiss/dasch-service-platform/issues/47)) ([d456715](https://www.github.com/dasch-swiss/dasch-service-platform/commit/d456715addc37eb1c3a464a22f614ee6570d3f3c))
* **metadata:** add beginnings of backend implementation in go (BS-28) ([#7](https://www.github.com/dasch-swiss/dasch-service-platform/issues/7)) ([7b601cf](https://www.github.com/dasch-swiss/dasch-service-platform/commit/7b601cf80608ff8cf681a6fd8e0244c10777826f))
* **metadata:** add organization and graphql ([#13](https://www.github.com/dasch-swiss/dasch-service-platform/issues/13)) ([378d182](https://www.github.com/dasch-swiss/dasch-service-platform/commit/378d182198a924098cdd0b69ab7f8f079be94cba))
* **project:** Project Service (DSP-1478) ([#28](https://www.github.com/dasch-swiss/dasch-service-platform/issues/28)) ([c84f745](https://www.github.com/dasch-swiss/dasch-service-platform/commit/c84f745fef59b9cb902a3bfb5aa613076b0fa621))
* **resource:** add resource service ([#8](https://www.github.com/dasch-swiss/dasch-service-platform/issues/8)) ([0a304eb](https://www.github.com/dasch-swiss/dasch-service-platform/commit/0a304ebf065e1628ecc0e81addad47a8ddde7a21))
