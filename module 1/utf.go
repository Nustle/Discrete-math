package main

func encode(utf32 []rune) []byte {
  var utf8 []byte
  n := len(utf32)
  for i := 0; i < n; i++ {
    if utf32[i] < (1 << 7) {
      utf8 = append(utf8, byte(utf32[i]))
    } else if utf32[i] < (1 << 11) { 
      second := utf32[i] & 0x3F
      first := (utf32[i] >> 6) & 0x1F
      utf8 = append(utf8, 0xC0|byte(first), 0x80|byte(second))
    } else if utf32[i] < (1 << 16) {
      third := utf32[i] & 0x3F
      second := (utf32[i] >> 6) & 0x3F
      first := (utf32[i] >> 12) & 0x0F
      utf8 = append(utf8, 0xE0|byte(first), 0x80|byte(second), 0x80|byte(third))
    } else {
      fourth := utf32[i] & 0x3F
      third := (utf32[i] >> 6) & 0x3F
      second := (utf32[i] >> 12) & 0x3F
      first := (utf32[i] >> 18) & 0x07
      utf8 = append(utf8, 0xF0|byte(first), 0x80|byte(second), 0x80|byte(third), 0x80|byte(fourth))
    }
  }
  return utf8
}

func decode(utf8 []byte) []rune {
  var utf32 []rune
  n := len(utf8)
  i := 0
  for i < n {
    if utf8[i] < 0x80 {
      utf32 = append(utf32, rune(utf8[i]))
      i += 1
    } else if utf8[i] < 0xE0 {
      first := rune(utf8[i]&0x1F) << 6
      second := rune(utf8[i+1] & 0x3F)
      utf32 = append(utf32, first|second)
      i += 2
    } else if utf8[i] < 0xF0 {
      first := rune(utf8[i]&0x0F) << 12
      second := rune(utf8[i+1]&0x3F) << 6
      third := rune(utf8[i+2] & 0x3F)
      utf32 = append(utf32, first|second|third)
      i += 3
    } else {
      first := rune(utf8[i]&0x07) << 18
      second := rune(utf8[i+1]&0x3F) << 12
      third := rune(utf8[i+2]&0x3F) << 6
      fourth := rune(utf8[i+3] & 0x3F)
      utf32 = append(utf32, first|second|third|fourth)
      i += 4
    }
  }
  return utf32
}

func main() {}
