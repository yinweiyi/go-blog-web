$(function () {
    var _$ = ["\x63\x6f\x6c\x6f\x72", '\x32', "\x72\x67\x62\x28", "\x2c", "\x2c", "\x29"];
    $["\x66\x6e"]["\x53\x74\x72\x65\x61\x6d\x65\x72"] = function (zh30e288558) {
        var zh305ab3d6d = {rgb: [0xaf, 0x48, 0x48], min: 0x48, max: 0xaf, step: 0x1, t: 0x14, properties: _$[0x0]};
        var zh3024447c6 = this;
        var zh30d7cbaca = $["\x65\x78\x74\x65\x6e\x64"]({}, zh305ab3d6d, zh30e288558);
        var zh30b086238 = _$[0x1];
        var zh30dc3cb99 = 0x1;
        var zh30589d39e = setInterval(function () {
            zh30d7cbaca["\x72\x67\x62"][zh30b086238] = zh30d7cbaca["\x72\x67\x62"][zh30b086238] + zh30dc3cb99 * zh30d7cbaca["\x73\x74\x65\x70"];
            if (zh30d7cbaca["\x72\x67\x62"][zh30b086238] >= zh30d7cbaca["\x6d\x61\x78"] || zh30d7cbaca["\x72\x67\x62"][zh30b086238] <= zh30d7cbaca["\x6d\x69\x6e"]) {
                if (zh30dc3cb99 == 0x1) {
                    zh30dc3cb99 = -0x1
                } else {
                    zh30dc3cb99 = 0x1
                }
                zh30b086238 = (zh30b086238 + 0x1) % 0x3
            }
            zh3024447c6["\x63\x73\x73"](zh30d7cbaca["\x70\x72\x6f\x70\x65\x72\x74\x69\x65\x73"], _$[0x2] + zh30d7cbaca["\x72\x67\x62"][0x0] + _$[0x3] + zh30d7cbaca["\x72\x67\x62"][0x1] + _$[0x4] + zh30d7cbaca["\x72\x67\x62"][0x2] + _$[0x5])
        }, zh30d7cbaca["\x74"])
    }

    $(".panel-heading").Streamer({max: 210, min: 50});
    $(".quote-headtips").Streamer({t: 10});
    $(".text img").addClass("img-responsive center-block");

})
