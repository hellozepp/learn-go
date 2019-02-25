pragma solidity ^ 0.4.18;

contract C {
    byte public a = 0x6c; // 0110 1100
    bytes1 public b = 0x6c; // 0110 1100
    bytes2 public c = 0x6c69; // 0110 1100 0110 1001
    bytes3 public d = 0x6c6979; // 0110 1100 0110 1001 0111 1001
    bytes4 public e = 0x6c697975; // 0110 1100 0110 1001 0111 1001 0111 0101

    // ...

    bytes8 public f = 0x6c69797565636875; // 0110 1100 0110 1001 0111 1001 0111 0101 0110 0101 0110 0011 0110 1000 0111 0101
    bytes9 public g = 0x6c697975656368756e; // // 0110 1100 0110 1001 0111 1001 0111 0101 0110 0101 0110 0011 0110 1000 0111 0101 0110 1110

    bytes9 public g = 0x6c697975656368756e;

    function gByteLength() constant returns (uint) {

        return g.length;
    }

    bytes public name = new bytes(1);//创建bytes字节数组
    // 0x6c697975656368756e

    // 0x6c697975656368756e
    // 初始化一个两个字节空间的字节数组 动态大小字节数组
    bytes public name = new bytes(2);

    // 设置字节数组的长度
    function setNameLength(uint len) {

        name.length = len;
        //你通过length方法将其长度修改为2时，字节数组中最后一个字节将被从字节数组中移除。
    }

    // 返回字节数组的长度
    function nameLength() constant returns (uint) {

        return name.length;
    }

    // 往字节数组中添加字节
    function pushAByte(byte b) {

        name.push(b);
    }

    function bytes32Content(bytes32 x) constant returns (bytes32) {
        // 0x6c697975656368756e000000000000000000000000000000000000000000000
        return x;
    }
    //固定大小字节数组转string，那么就需要先将字节数组转动态字节数组，再转字符串
    function bytes32ToString(bytes32 x) constant returns (string) {
        //0x6c697975656368756e0000000000000000000000000000000000000000000000
        bytes memory bytesString = new bytes(32);
        uint charCount = 0;
        for (uint j = 0; j < 32; j++) {
            // 6 * 2 12
            // 0000001 1000
            // 0x697975656368756e0000000000000000000000000000000000000000000
            // byte char = byte(bytes32(uint(x) * 2 ** (8 * j)));
            byte char = byte(bytes32(uint(x) << (8 * j)));
            //左移j
            if (char != 0) {
                bytesString[charCount] = char;
                // 0 0x6c
                // 1 0x69
                // 2 0x79
                // 3 0x95
                // ......
                // 8 0x6e
                charCount++;
            }
        }
        bytes memory bytesStringTrimmed = new bytes(charCount);
        //0x6c697975
        for (j = 0; j < charCount; j++) {
            bytesStringTrimmed[j] = bytesString[j];
        }
        return string(bytesStringTrimmed);
    }
    //固定大小字节数组转动态大小字节数组正确的姿势
    bytes9 name9 = 0x6c697975656368756e;

    function fixedSizeByteArraysToDynamicallySizedByteArray() constant returns (bytes) {

        bytes memory names = new bytes(name9.length);

        for (uint i = 0; i < name9.length; i++) {

            names[i] = name9[i];
        }

        return names;
    }
    //固定大小字节数组（Fixed-size byte arrays）不能直接转换为string,长度和字节均不可修改,可以通过索引读取(只读)对应索引的字节
    bytes9 names = 0x6c697975656368756e;

    function namesToString() constant returns (string) {

        return string(names);
    }
    // 0x6c

    function uintValue() constant returns (uint) {

        return uint(0x6c);
    }

    function bytes32To0x6c() constant returns (bytes32) {

        return bytes32(0x6c);
    }

    function bytes32To0x6cLeft00() constant returns (bytes32) {

        return bytes32(uint(0x6c) * 2 ** (8 * 0));
    }

    function bytes32To0x6cLeft01() constant returns (bytes32) {

        return bytes32(uint(0x6c) * 2 ** (8 * 1));
        //左移1位
    }

    function bytes32To0x6cLeft31() constant returns (bytes32) {

        return bytes32(uint(0x6c) * 2 ** (8 * 31));
    }
    //bytes0 ~ bytes32直接声明的不可变字节数组中，长度不可变，内容不可修改。而byte[len] b创建的字节数组中，长度不可变，但是内容可修改
    bytes9 a = 0x6c697975656368756e;
    byte[9] aa = [byte(0x6c), 0x69, 0x79, 0x75, 0x65, 0x63, 0x68, 0x75, 0x6e];//固定长度的数组（Arrays）
    function pushUintToAa() public {

        aa.push(6);
        //我们可以通过索引修改里面的值，但是不可修改数组长度以及不可通过push添加存储内容
    }

    byte[] cc = new byte[](10);

    function setAIndex0Byte() public {
        // 错误，不可修改
        a[0] = 0x89;
    }

    function setAAIndex0Byte() public {

        aa[0] = 0x89;
    }

    function setCC() public {

        for (uint i = 0; i < a.length; i++) {

            cc.push(a[i]);
        }
    }
    //    变长数组，可随时通过length修改它的长度
    uint [] T = [1, 2, 3, 4, 5];//storage类型的可变数组

    function T_Length() constant public returns (uint) {

        return T.length;
    }

    function setTLength(uint len) public {

        T.length = len;
    }

    function pushUintToT() public {

        T.push(6);
    }

    function numbers() constant public returns (uint) {
        uint num = 0;
        for (uint i = 0; i < T.length; i++) {
            num = num + T[i];
        }
        return num;
    }

    //二维数组
    uint [2][3] T = [[1, 2], [3, 4], [5, 6]];

    function T_len() constant public returns (uint) {

        return T.length;
        // 3
    }
    //    上面的数组T是storage类型的数组，对于storage类型的数组，数组里面可以存放任意类型的值（比如：其它数组，结构体，字典／映射等等）。
    //对于memory类型的数组，如果它是一个public类型的函数的参数，那么它里面的内容不能是一个mapping(映射／字典)，并且它必须是一个ABI类型。

    function f(uint len) {
        //memory数组一旦创建，它不可通过length修改其长度
        uint[] memory a = new uint[](7);
        bytes memory b = new bytes(len);
        // 在这段代码中 a.length == 7 、b.length == len
        a[6] = 8;
    }
    //数组字面量 Array Literals
    function f() public {
        g([uint(1), 2, 3]);//[1, 2, 3]是 uint8[3] memory
    }

    function g(uint[3] _data) public {
        // ...
    }

    function f1() public {
//memory类型的固定长度的数组不可直接赋值给storage/memory类型的可变数组
        uint[3] memory x = [uint(1), 3, 4];
    }
}