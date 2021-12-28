// Generated from Ruba.g4 by ANTLR 4.7.

package parser // Ruba

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 306,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 33, 10, 3, 12, 3, 14,
	3, 36, 11, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 54, 10, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 60, 10, 10, 3, 10, 5, 10, 63, 10, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 7, 10, 71, 10, 10, 12, 10, 14, 10, 74, 11, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7,
	10, 86, 10, 10, 12, 10, 14, 10, 89, 11, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	94, 10, 10, 3, 10, 3, 10, 5, 10, 98, 10, 10, 5, 10, 100, 10, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 111, 10,
	10, 12, 10, 14, 10, 114, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	5, 10, 132, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 146, 10, 10, 12, 10, 14, 10, 149,
	11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 156, 10, 10, 3, 10, 3,
	10, 5, 10, 160, 10, 10, 3, 10, 3, 10, 5, 10, 164, 10, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 5, 11, 170, 10, 11, 3, 12, 3, 12, 5, 12, 174, 10, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 7, 12, 180, 10, 12, 12, 12, 14, 12, 183, 11, 12,
	5, 12, 185, 10, 12, 3, 12, 3, 12, 5, 12, 189, 10, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 5, 13, 196, 10, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14,
	223, 10, 14, 12, 14, 14, 14, 226, 11, 14, 5, 14, 228, 10, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 7, 14, 235, 10, 14, 12, 14, 14, 14, 238, 11, 14,
	3, 14, 5, 14, 241, 10, 14, 5, 14, 243, 10, 14, 3, 14, 5, 14, 246, 10, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 7, 14, 283, 10, 14, 12, 14, 14, 14, 286, 11, 14,
	5, 14, 288, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 7, 14, 301, 10, 14, 12, 14, 14, 14, 304, 11, 14,
	3, 14, 2, 3, 26, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2,
	6, 3, 2, 3, 6, 3, 2, 30, 31, 3, 2, 5, 6, 3, 2, 3, 4, 2, 353, 2, 28, 3,
	2, 2, 2, 4, 34, 3, 2, 2, 2, 6, 37, 3, 2, 2, 2, 8, 39, 3, 2, 2, 2, 10, 41,
	3, 2, 2, 2, 12, 43, 3, 2, 2, 2, 14, 45, 3, 2, 2, 2, 16, 53, 3, 2, 2, 2,
	18, 163, 3, 2, 2, 2, 20, 169, 3, 2, 2, 2, 22, 173, 3, 2, 2, 2, 24, 195,
	3, 2, 2, 2, 26, 245, 3, 2, 2, 2, 28, 29, 5, 4, 3, 2, 29, 30, 7, 2, 2, 3,
	30, 3, 3, 2, 2, 2, 31, 33, 5, 18, 10, 2, 32, 31, 3, 2, 2, 2, 33, 36, 3,
	2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 5, 3, 2, 2, 2, 36,
	34, 3, 2, 2, 2, 37, 38, 7, 43, 2, 2, 38, 7, 3, 2, 2, 2, 39, 40, 7, 44,
	2, 2, 40, 9, 3, 2, 2, 2, 41, 42, 7, 46, 2, 2, 42, 11, 3, 2, 2, 2, 43, 44,
	7, 45, 2, 2, 44, 13, 3, 2, 2, 2, 45, 46, 9, 2, 2, 2, 46, 15, 3, 2, 2, 2,
	47, 48, 7, 7, 2, 2, 48, 54, 5, 6, 4, 2, 49, 50, 7, 8, 2, 2, 50, 51, 5,
	26, 14, 2, 51, 52, 7, 9, 2, 2, 52, 54, 3, 2, 2, 2, 53, 47, 3, 2, 2, 2,
	53, 49, 3, 2, 2, 2, 54, 17, 3, 2, 2, 2, 55, 56, 7, 10, 2, 2, 56, 59, 5,
	12, 7, 2, 57, 58, 7, 11, 2, 2, 58, 60, 5, 6, 4, 2, 59, 57, 3, 2, 2, 2,
	59, 60, 3, 2, 2, 2, 60, 164, 3, 2, 2, 2, 61, 63, 7, 12, 2, 2, 62, 61, 3,
	2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 65, 5, 6, 4, 2, 65,
	66, 7, 13, 2, 2, 66, 67, 5, 26, 14, 2, 67, 164, 3, 2, 2, 2, 68, 72, 5,
	6, 4, 2, 69, 71, 5, 16, 9, 2, 70, 69, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72,
	70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 72, 3, 2, 2,
	2, 75, 76, 7, 14, 2, 2, 76, 77, 5, 26, 14, 2, 77, 164, 3, 2, 2, 2, 78,
	164, 5, 22, 12, 2, 79, 80, 7, 15, 2, 2, 80, 81, 5, 6, 4, 2, 81, 99, 7,
	16, 2, 2, 82, 87, 5, 6, 4, 2, 83, 84, 7, 17, 2, 2, 84, 86, 5, 6, 4, 2,
	85, 83, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3,
	2, 2, 2, 88, 93, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 91, 7, 17, 2, 2, 91,
	92, 7, 18, 2, 2, 92, 94, 5, 6, 4, 2, 93, 90, 3, 2, 2, 2, 93, 94, 3, 2,
	2, 2, 94, 98, 3, 2, 2, 2, 95, 96, 7, 18, 2, 2, 96, 98, 5, 6, 4, 2, 97,
	82, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 100, 3, 2, 2, 2, 99, 97, 3, 2,
	2, 2, 99, 100, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 7, 19, 2, 2,
	102, 103, 7, 20, 2, 2, 103, 104, 5, 4, 3, 2, 104, 105, 7, 21, 2, 2, 105,
	164, 3, 2, 2, 2, 106, 107, 7, 22, 2, 2, 107, 112, 5, 6, 4, 2, 108, 109,
	7, 17, 2, 2, 109, 111, 5, 6, 4, 2, 110, 108, 3, 2, 2, 2, 111, 114, 3, 2,
	2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 115, 3, 2, 2, 2,
	114, 112, 3, 2, 2, 2, 115, 116, 7, 23, 2, 2, 116, 117, 5, 26, 14, 2, 117,
	118, 7, 20, 2, 2, 118, 119, 5, 4, 3, 2, 119, 120, 7, 21, 2, 2, 120, 164,
	3, 2, 2, 2, 121, 122, 7, 22, 2, 2, 122, 123, 5, 26, 14, 2, 123, 124, 7,
	20, 2, 2, 124, 125, 5, 4, 3, 2, 125, 131, 7, 21, 2, 2, 126, 127, 7, 24,
	2, 2, 127, 128, 7, 20, 2, 2, 128, 129, 5, 4, 3, 2, 129, 130, 7, 21, 2,
	2, 130, 132, 3, 2, 2, 2, 131, 126, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132,
	164, 3, 2, 2, 2, 133, 134, 7, 25, 2, 2, 134, 135, 5, 26, 14, 2, 135, 136,
	7, 20, 2, 2, 136, 137, 5, 4, 3, 2, 137, 147, 7, 21, 2, 2, 138, 139, 7,
	24, 2, 2, 139, 140, 7, 25, 2, 2, 140, 141, 5, 26, 14, 2, 141, 142, 7, 20,
	2, 2, 142, 143, 5, 4, 3, 2, 143, 144, 7, 21, 2, 2, 144, 146, 3, 2, 2, 2,
	145, 138, 3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147,
	148, 3, 2, 2, 2, 148, 155, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 151,
	7, 24, 2, 2, 151, 152, 7, 20, 2, 2, 152, 153, 5, 4, 3, 2, 153, 154, 7,
	21, 2, 2, 154, 156, 3, 2, 2, 2, 155, 150, 3, 2, 2, 2, 155, 156, 3, 2, 2,
	2, 156, 164, 3, 2, 2, 2, 157, 159, 7, 26, 2, 2, 158, 160, 5, 26, 14, 2,
	159, 158, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 164, 3, 2, 2, 2, 161,
	164, 7, 27, 2, 2, 162, 164, 5, 26, 14, 2, 163, 55, 3, 2, 2, 2, 163, 62,
	3, 2, 2, 2, 163, 68, 3, 2, 2, 2, 163, 78, 3, 2, 2, 2, 163, 79, 3, 2, 2,
	2, 163, 106, 3, 2, 2, 2, 163, 121, 3, 2, 2, 2, 163, 133, 3, 2, 2, 2, 163,
	157, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 162, 3, 2, 2, 2, 164, 19, 3,
	2, 2, 2, 165, 170, 5, 26, 14, 2, 166, 167, 5, 26, 14, 2, 167, 168, 7, 18,
	2, 2, 168, 170, 3, 2, 2, 2, 169, 165, 3, 2, 2, 2, 169, 166, 3, 2, 2, 2,
	170, 21, 3, 2, 2, 2, 171, 174, 5, 6, 4, 2, 172, 174, 5, 26, 14, 2, 173,
	171, 3, 2, 2, 2, 173, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 184,
	7, 16, 2, 2, 176, 181, 5, 20, 11, 2, 177, 178, 7, 17, 2, 2, 178, 180, 5,
	20, 11, 2, 179, 177, 3, 2, 2, 2, 180, 183, 3, 2, 2, 2, 181, 179, 3, 2,
	2, 2, 181, 182, 3, 2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2,
	184, 176, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186,
	188, 7, 19, 2, 2, 187, 189, 7, 28, 2, 2, 188, 187, 3, 2, 2, 2, 188, 189,
	3, 2, 2, 2, 189, 23, 3, 2, 2, 2, 190, 196, 5, 6, 4, 2, 191, 192, 7, 8,
	2, 2, 192, 193, 5, 26, 14, 2, 193, 194, 7, 9, 2, 2, 194, 196, 3, 2, 2,
	2, 195, 190, 3, 2, 2, 2, 195, 191, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197,
	198, 7, 29, 2, 2, 198, 199, 5, 26, 14, 2, 199, 25, 3, 2, 2, 2, 200, 201,
	8, 14, 1, 2, 201, 202, 7, 16, 2, 2, 202, 203, 5, 26, 14, 2, 203, 204, 7,
	19, 2, 2, 204, 246, 3, 2, 2, 2, 205, 246, 9, 3, 2, 2, 206, 246, 7, 32,
	2, 2, 207, 246, 5, 6, 4, 2, 208, 246, 5, 8, 5, 2, 209, 246, 5, 10, 6, 2,
	210, 246, 5, 12, 7, 2, 211, 212, 7, 33, 2, 2, 212, 246, 5, 26, 14, 17,
	213, 214, 7, 4, 2, 2, 214, 246, 5, 26, 14, 16, 215, 216, 5, 6, 4, 2, 216,
	217, 7, 41, 2, 2, 217, 246, 3, 2, 2, 2, 218, 227, 7, 8, 2, 2, 219, 224,
	5, 26, 14, 2, 220, 221, 7, 17, 2, 2, 221, 223, 5, 26, 14, 2, 222, 220,
	3, 2, 2, 2, 223, 226, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 224, 225, 3, 2,
	2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 227, 219, 3, 2, 2, 2,
	227, 228, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 246, 7, 9, 2, 2, 230,
	242, 7, 20, 2, 2, 231, 236, 5, 24, 13, 2, 232, 233, 7, 17, 2, 2, 233, 235,
	5, 24, 13, 2, 234, 232, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3,
	2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 240, 3, 2, 2, 2, 238, 236, 3, 2, 2,
	2, 239, 241, 7, 17, 2, 2, 240, 239, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241,
	243, 3, 2, 2, 2, 242, 231, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 244,
	3, 2, 2, 2, 244, 246, 7, 21, 2, 2, 245, 200, 3, 2, 2, 2, 245, 205, 3, 2,
	2, 2, 245, 206, 3, 2, 2, 2, 245, 207, 3, 2, 2, 2, 245, 208, 3, 2, 2, 2,
	245, 209, 3, 2, 2, 2, 245, 210, 3, 2, 2, 2, 245, 211, 3, 2, 2, 2, 245,
	213, 3, 2, 2, 2, 245, 215, 3, 2, 2, 2, 245, 218, 3, 2, 2, 2, 245, 230,
	3, 2, 2, 2, 246, 302, 3, 2, 2, 2, 247, 248, 12, 15, 2, 2, 248, 249, 9,
	4, 2, 2, 249, 301, 5, 26, 14, 16, 250, 251, 12, 14, 2, 2, 251, 252, 9,
	5, 2, 2, 252, 301, 5, 26, 14, 15, 253, 254, 12, 13, 2, 2, 254, 255, 7,
	34, 2, 2, 255, 301, 5, 26, 14, 14, 256, 257, 12, 12, 2, 2, 257, 258, 7,
	35, 2, 2, 258, 301, 5, 26, 14, 13, 259, 260, 12, 11, 2, 2, 260, 261, 7,
	36, 2, 2, 261, 301, 5, 26, 14, 12, 262, 263, 12, 10, 2, 2, 263, 264, 7,
	37, 2, 2, 264, 301, 5, 26, 14, 11, 265, 266, 12, 9, 2, 2, 266, 267, 7,
	38, 2, 2, 267, 301, 5, 26, 14, 10, 268, 269, 12, 8, 2, 2, 269, 270, 7,
	39, 2, 2, 270, 301, 5, 26, 14, 9, 271, 272, 12, 3, 2, 2, 272, 273, 7, 42,
	2, 2, 273, 274, 5, 26, 14, 2, 274, 275, 7, 29, 2, 2, 275, 276, 5, 26, 14,
	4, 276, 301, 3, 2, 2, 2, 277, 278, 12, 24, 2, 2, 278, 287, 7, 16, 2, 2,
	279, 284, 5, 20, 11, 2, 280, 281, 7, 17, 2, 2, 281, 283, 5, 20, 11, 2,
	282, 280, 3, 2, 2, 2, 283, 286, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 284,
	285, 3, 2, 2, 2, 285, 288, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 287, 279,
	3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 301, 7, 19,
	2, 2, 290, 291, 12, 22, 2, 2, 291, 292, 7, 8, 2, 2, 292, 293, 5, 26, 14,
	2, 293, 294, 7, 9, 2, 2, 294, 301, 3, 2, 2, 2, 295, 296, 12, 21, 2, 2,
	296, 297, 7, 7, 2, 2, 297, 301, 5, 6, 4, 2, 298, 299, 12, 7, 2, 2, 299,
	301, 7, 40, 2, 2, 300, 247, 3, 2, 2, 2, 300, 250, 3, 2, 2, 2, 300, 253,
	3, 2, 2, 2, 300, 256, 3, 2, 2, 2, 300, 259, 3, 2, 2, 2, 300, 262, 3, 2,
	2, 2, 300, 265, 3, 2, 2, 2, 300, 268, 3, 2, 2, 2, 300, 271, 3, 2, 2, 2,
	300, 277, 3, 2, 2, 2, 300, 290, 3, 2, 2, 2, 300, 295, 3, 2, 2, 2, 300,
	298, 3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302, 303,
	3, 2, 2, 2, 303, 27, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 33, 34, 53, 59,
	62, 72, 87, 93, 97, 99, 112, 131, 147, 155, 159, 163, 169, 173, 181, 184,
	188, 195, 224, 227, 236, 240, 242, 245, 284, 287, 300, 302,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "'/'", "'*'", "'.'", "'['", "']'", "'import'", "'as'",
	"'pub'", "':='", "'='", "'func'", "'('", "','", "'...'", "')'", "'{'",
	"'}'", "'for'", "'in'", "'else'", "'if'", "'return'", "'break'", "';'",
	"':'", "'true'", "'false'", "'null'", "'!'", "'=='", "'!='", "'>'", "'<='",
	"'<'", "'>='", "'++'", "'--'", "'?'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "Ident", "Int", "Str", "Float", "WS",
}

var ruleNames = []string{
	"program", "block", "ident", "intType", "floatType", "stringType", "mathSign",
	"nestable", "stmt", "fnArg", "fnCall", "objectField", "expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RubaParser struct {
	*antlr.BaseParser
}

func NewRubaParser(input antlr.TokenStream) *RubaParser {
	this := new(RubaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Ruba.g4"

	return this
}

// RubaParser tokens.
const (
	RubaParserEOF   = antlr.TokenEOF
	RubaParserT__0  = 1
	RubaParserT__1  = 2
	RubaParserT__2  = 3
	RubaParserT__3  = 4
	RubaParserT__4  = 5
	RubaParserT__5  = 6
	RubaParserT__6  = 7
	RubaParserT__7  = 8
	RubaParserT__8  = 9
	RubaParserT__9  = 10
	RubaParserT__10 = 11
	RubaParserT__11 = 12
	RubaParserT__12 = 13
	RubaParserT__13 = 14
	RubaParserT__14 = 15
	RubaParserT__15 = 16
	RubaParserT__16 = 17
	RubaParserT__17 = 18
	RubaParserT__18 = 19
	RubaParserT__19 = 20
	RubaParserT__20 = 21
	RubaParserT__21 = 22
	RubaParserT__22 = 23
	RubaParserT__23 = 24
	RubaParserT__24 = 25
	RubaParserT__25 = 26
	RubaParserT__26 = 27
	RubaParserT__27 = 28
	RubaParserT__28 = 29
	RubaParserT__29 = 30
	RubaParserT__30 = 31
	RubaParserT__31 = 32
	RubaParserT__32 = 33
	RubaParserT__33 = 34
	RubaParserT__34 = 35
	RubaParserT__35 = 36
	RubaParserT__36 = 37
	RubaParserT__37 = 38
	RubaParserT__38 = 39
	RubaParserT__39 = 40
	RubaParserIdent = 41
	RubaParserInt   = 42
	RubaParserStr   = 43
	RubaParserFloat = 44
	RubaParserWS    = 45
)

// RubaParser rules.
const (
	RubaParserRULE_program     = 0
	RubaParserRULE_block       = 1
	RubaParserRULE_ident       = 2
	RubaParserRULE_intType     = 3
	RubaParserRULE_floatType   = 4
	RubaParserRULE_stringType  = 5
	RubaParserRULE_mathSign    = 6
	RubaParserRULE_nestable    = 7
	RubaParserRULE_stmt        = 8
	RubaParserRULE_fnArg       = 9
	RubaParserRULE_fnCall      = 10
	RubaParserRULE_objectField = 11
	RubaParserRULE_expr        = 12
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(RubaParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RubaParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Block()
	}
	{
		p.SetState(27)
		p.Match(RubaParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RubaParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RubaParserT__1)|(1<<RubaParserT__5)|(1<<RubaParserT__7)|(1<<RubaParserT__9)|(1<<RubaParserT__12)|(1<<RubaParserT__13)|(1<<RubaParserT__17)|(1<<RubaParserT__19)|(1<<RubaParserT__22)|(1<<RubaParserT__23)|(1<<RubaParserT__24)|(1<<RubaParserT__27)|(1<<RubaParserT__28)|(1<<RubaParserT__29)|(1<<RubaParserT__30))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(RubaParserIdent-41))|(1<<(RubaParserInt-41))|(1<<(RubaParserStr-41))|(1<<(RubaParserFloat-41)))) != 0) {
		{
			p.SetState(29)
			p.Stmt()
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) Ident() antlr.TerminalNode {
	return s.GetToken(RubaParserIdent, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RubaParserRULE_ident)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(RubaParserIdent)
	}

	return localctx
}

// IIntTypeContext is an interface to support dynamic dispatch.
type IIntTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntTypeContext differentiates from other interfaces.
	IsIntTypeContext()
}

type IntTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntTypeContext() *IntTypeContext {
	var p = new(IntTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_intType
	return p
}

func (*IntTypeContext) IsIntTypeContext() {}

func NewIntTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntTypeContext {
	var p = new(IntTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_intType

	return p
}

func (s *IntTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IntTypeContext) Int() antlr.TerminalNode {
	return s.GetToken(RubaParserInt, 0)
}

func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterIntType(s)
	}
}

func (s *IntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitIntType(s)
	}
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) IntType() (localctx IIntTypeContext) {
	localctx = NewIntTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RubaParserRULE_intType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(RubaParserInt)
	}

	return localctx
}

// IFloatTypeContext is an interface to support dynamic dispatch.
type IFloatTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatTypeContext differentiates from other interfaces.
	IsFloatTypeContext()
}

type FloatTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatTypeContext() *FloatTypeContext {
	var p = new(FloatTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_floatType
	return p
}

func (*FloatTypeContext) IsFloatTypeContext() {}

func NewFloatTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatTypeContext {
	var p = new(FloatTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_floatType

	return p
}

func (s *FloatTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatTypeContext) Float() antlr.TerminalNode {
	return s.GetToken(RubaParserFloat, 0)
}

func (s *FloatTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterFloatType(s)
	}
}

func (s *FloatTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitFloatType(s)
	}
}

func (s *FloatTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitFloatType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) FloatType() (localctx IFloatTypeContext) {
	localctx = NewFloatTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RubaParserRULE_floatType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(RubaParserFloat)
	}

	return localctx
}

// IStringTypeContext is an interface to support dynamic dispatch.
type IStringTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringTypeContext differentiates from other interfaces.
	IsStringTypeContext()
}

type StringTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringTypeContext() *StringTypeContext {
	var p = new(StringTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_stringType
	return p
}

func (*StringTypeContext) IsStringTypeContext() {}

func NewStringTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringTypeContext {
	var p = new(StringTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_stringType

	return p
}

func (s *StringTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StringTypeContext) Str() antlr.TerminalNode {
	return s.GetToken(RubaParserStr, 0)
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterStringType(s)
	}
}

func (s *StringTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitStringType(s)
	}
}

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) StringType() (localctx IStringTypeContext) {
	localctx = NewStringTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RubaParserRULE_stringType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(RubaParserStr)
	}

	return localctx
}

// IMathSignContext is an interface to support dynamic dispatch.
type IMathSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathSignContext differentiates from other interfaces.
	IsMathSignContext()
}

type MathSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathSignContext() *MathSignContext {
	var p = new(MathSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_mathSign
	return p
}

func (*MathSignContext) IsMathSignContext() {}

func NewMathSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathSignContext {
	var p = new(MathSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_mathSign

	return p
}

func (s *MathSignContext) GetParser() antlr.Parser { return s.parser }
func (s *MathSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterMathSign(s)
	}
}

func (s *MathSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitMathSign(s)
	}
}

func (s *MathSignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitMathSign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) MathSign() (localctx IMathSignContext) {
	localctx = NewMathSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RubaParserRULE_mathSign)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RubaParserT__0)|(1<<RubaParserT__1)|(1<<RubaParserT__2)|(1<<RubaParserT__3))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// INestableContext is an interface to support dynamic dispatch.
type INestableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDOT returns the DOT rule contexts.
	GetDOT() IIdentContext

	// GetINDEX returns the INDEX rule contexts.
	GetINDEX() IExprContext

	// SetDOT sets the DOT rule contexts.
	SetDOT(IIdentContext)

	// SetINDEX sets the INDEX rule contexts.
	SetINDEX(IExprContext)

	// IsNestableContext differentiates from other interfaces.
	IsNestableContext()
}

type NestableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	DOT    IIdentContext
	INDEX  IExprContext
}

func NewEmptyNestableContext() *NestableContext {
	var p = new(NestableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_nestable
	return p
}

func (*NestableContext) IsNestableContext() {}

func NewNestableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestableContext {
	var p = new(NestableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_nestable

	return p
}

func (s *NestableContext) GetParser() antlr.Parser { return s.parser }

func (s *NestableContext) GetDOT() IIdentContext { return s.DOT }

func (s *NestableContext) GetINDEX() IExprContext { return s.INDEX }

func (s *NestableContext) SetDOT(v IIdentContext) { s.DOT = v }

func (s *NestableContext) SetINDEX(v IExprContext) { s.INDEX = v }

func (s *NestableContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *NestableContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NestableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterNestable(s)
	}
}

func (s *NestableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitNestable(s)
	}
}

func (s *NestableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitNestable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) Nestable() (localctx INestableContext) {
	localctx = NewNestableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RubaParserRULE_nestable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RubaParserT__4:
		{
			p.SetState(45)
			p.Match(RubaParserT__4)
		}
		{
			p.SetState(46)

			var _x = p.Ident()

			localctx.(*NestableContext).DOT = _x
		}

	case RubaParserT__5:
		{
			p.SetState(47)
			p.Match(RubaParserT__5)
		}
		{
			p.SetState(48)

			var _x = p.expr(0)

			localctx.(*NestableContext).INDEX = _x
		}
		{
			p.SetState(49)
			p.Match(RubaParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyFrom(ctx *StmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForInStmtContext struct {
	*StmtContext
	_ident IIdentContext
	NAMES  []IIdentContext
	TARGET IExprContext
	BODY   IBlockContext
}

func NewForInStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInStmtContext {
	var p = new(ForInStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ForInStmtContext) Get_ident() IIdentContext { return s._ident }

func (s *ForInStmtContext) GetTARGET() IExprContext { return s.TARGET }

func (s *ForInStmtContext) GetBODY() IBlockContext { return s.BODY }

func (s *ForInStmtContext) Set_ident(v IIdentContext) { s._ident = v }

func (s *ForInStmtContext) SetTARGET(v IExprContext) { s.TARGET = v }

func (s *ForInStmtContext) SetBODY(v IBlockContext) { s.BODY = v }

func (s *ForInStmtContext) GetNAMES() []IIdentContext { return s.NAMES }

func (s *ForInStmtContext) SetNAMES(v []IIdentContext) { s.NAMES = v }

func (s *ForInStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInStmtContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *ForInStmtContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ForInStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForInStmtContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForInStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterForInStmt(s)
	}
}

func (s *ForInStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitForInStmt(s)
	}
}

func (s *ForInStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitForInStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type FnCallStmtContext struct {
	*StmtContext
}

func NewFnCallStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FnCallStmtContext {
	var p = new(FnCallStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *FnCallStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnCallStmtContext) FnCall() IFnCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnCallContext)
}

func (s *FnCallStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterFnCallStmt(s)
	}
}

func (s *FnCallStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitFnCallStmt(s)
	}
}

func (s *FnCallStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitFnCallStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImportStmtContext struct {
	*StmtContext
	NAME  IStringTypeContext
	ALIAS IIdentContext
}

func NewImportStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportStmtContext {
	var p = new(ImportStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ImportStmtContext) GetNAME() IStringTypeContext { return s.NAME }

func (s *ImportStmtContext) GetALIAS() IIdentContext { return s.ALIAS }

func (s *ImportStmtContext) SetNAME(v IStringTypeContext) { s.NAME = v }

func (s *ImportStmtContext) SetALIAS(v IIdentContext) { s.ALIAS = v }

func (s *ImportStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStmtContext) StringType() IStringTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringTypeContext)
}

func (s *ImportStmtContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ImportStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterImportStmt(s)
	}
}

func (s *ImportStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitImportStmt(s)
	}
}

func (s *ImportStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitImportStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStmtContext struct {
	*StmtContext
	COND     IExprContext
	BODY     IBlockContext
	_expr    IExprContext
	ELIFCOND []IExprContext
	_block   IBlockContext
	ELIFBODY []IBlockContext
	ELSEBODY IBlockContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *IfStmtContext) GetCOND() IExprContext { return s.COND }

func (s *IfStmtContext) GetBODY() IBlockContext { return s.BODY }

func (s *IfStmtContext) Get_expr() IExprContext { return s._expr }

func (s *IfStmtContext) Get_block() IBlockContext { return s._block }

func (s *IfStmtContext) GetELSEBODY() IBlockContext { return s.ELSEBODY }

func (s *IfStmtContext) SetCOND(v IExprContext) { s.COND = v }

func (s *IfStmtContext) SetBODY(v IBlockContext) { s.BODY = v }

func (s *IfStmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfStmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *IfStmtContext) SetELSEBODY(v IBlockContext) { s.ELSEBODY = v }

func (s *IfStmtContext) GetELIFCOND() []IExprContext { return s.ELIFCOND }

func (s *IfStmtContext) GetELIFBODY() []IBlockContext { return s.ELIFBODY }

func (s *IfStmtContext) SetELIFCOND(v []IExprContext) { s.ELIFCOND = v }

func (s *IfStmtContext) SetELIFBODY(v []IBlockContext) { s.ELIFBODY = v }

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IfStmtContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprStmtContext struct {
	*StmtContext
	E IExprContext
}

func NewExprStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStmtContext {
	var p = new(ExprStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ExprStmtContext) GetE() IExprContext { return s.E }

func (s *ExprStmtContext) SetE(v IExprContext) { s.E = v }

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignStmtContext struct {
	*StmtContext
	ROOT      IIdentContext
	_nestable INestableContext
	CHAIN     []INestableContext
	VALUE     IExprContext
}

func NewAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStmtContext {
	var p = new(AssignStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *AssignStmtContext) GetROOT() IIdentContext { return s.ROOT }

func (s *AssignStmtContext) Get_nestable() INestableContext { return s._nestable }

func (s *AssignStmtContext) GetVALUE() IExprContext { return s.VALUE }

func (s *AssignStmtContext) SetROOT(v IIdentContext) { s.ROOT = v }

func (s *AssignStmtContext) Set_nestable(v INestableContext) { s._nestable = v }

func (s *AssignStmtContext) SetVALUE(v IExprContext) { s.VALUE = v }

func (s *AssignStmtContext) GetCHAIN() []INestableContext { return s.CHAIN }

func (s *AssignStmtContext) SetCHAIN(v []INestableContext) { s.CHAIN = v }

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *AssignStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignStmtContext) AllNestable() []INestableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INestableContext)(nil)).Elem())
	var tst = make([]INestableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INestableContext)
		}
	}

	return tst
}

func (s *AssignStmtContext) Nestable(i int) INestableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INestableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INestableContext)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	*StmtContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclareStmtContext struct {
	*StmtContext
	PUB   antlr.Token
	NAME  IIdentContext
	VALUE IExprContext
}

func NewDeclareStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclareStmtContext {
	var p = new(DeclareStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *DeclareStmtContext) GetPUB() antlr.Token { return s.PUB }

func (s *DeclareStmtContext) SetPUB(v antlr.Token) { s.PUB = v }

func (s *DeclareStmtContext) GetNAME() IIdentContext { return s.NAME }

func (s *DeclareStmtContext) GetVALUE() IExprContext { return s.VALUE }

func (s *DeclareStmtContext) SetNAME(v IIdentContext) { s.NAME = v }

func (s *DeclareStmtContext) SetVALUE(v IExprContext) { s.VALUE = v }

func (s *DeclareStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclareStmtContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *DeclareStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclareStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterDeclareStmt(s)
	}
}

func (s *DeclareStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitDeclareStmt(s)
	}
}

func (s *DeclareStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitDeclareStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type FnDeclareStmtContext struct {
	*StmtContext
	NAME    IIdentContext
	_ident  IIdentContext
	ARGS    []IIdentContext
	RESTARG IIdentContext
	BODY    IBlockContext
}

func NewFnDeclareStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FnDeclareStmtContext {
	var p = new(FnDeclareStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *FnDeclareStmtContext) GetNAME() IIdentContext { return s.NAME }

func (s *FnDeclareStmtContext) Get_ident() IIdentContext { return s._ident }

func (s *FnDeclareStmtContext) GetRESTARG() IIdentContext { return s.RESTARG }

func (s *FnDeclareStmtContext) GetBODY() IBlockContext { return s.BODY }

func (s *FnDeclareStmtContext) SetNAME(v IIdentContext) { s.NAME = v }

func (s *FnDeclareStmtContext) Set_ident(v IIdentContext) { s._ident = v }

func (s *FnDeclareStmtContext) SetRESTARG(v IIdentContext) { s.RESTARG = v }

func (s *FnDeclareStmtContext) SetBODY(v IBlockContext) { s.BODY = v }

func (s *FnDeclareStmtContext) GetARGS() []IIdentContext { return s.ARGS }

func (s *FnDeclareStmtContext) SetARGS(v []IIdentContext) { s.ARGS = v }

func (s *FnDeclareStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnDeclareStmtContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *FnDeclareStmtContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FnDeclareStmtContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FnDeclareStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterFnDeclareStmt(s)
	}
}

func (s *FnDeclareStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitFnDeclareStmt(s)
	}
}

func (s *FnDeclareStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitFnDeclareStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForCondStmtContext struct {
	*StmtContext
	COND     IExprContext
	BODY     IBlockContext
	ELSEBODY IBlockContext
}

func NewForCondStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForCondStmtContext {
	var p = new(ForCondStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ForCondStmtContext) GetCOND() IExprContext { return s.COND }

func (s *ForCondStmtContext) GetBODY() IBlockContext { return s.BODY }

func (s *ForCondStmtContext) GetELSEBODY() IBlockContext { return s.ELSEBODY }

func (s *ForCondStmtContext) SetCOND(v IExprContext) { s.COND = v }

func (s *ForCondStmtContext) SetBODY(v IBlockContext) { s.BODY = v }

func (s *ForCondStmtContext) SetELSEBODY(v IBlockContext) { s.ELSEBODY = v }

func (s *ForCondStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForCondStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForCondStmtContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *ForCondStmtContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForCondStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterForCondStmt(s)
	}
}

func (s *ForCondStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitForCondStmt(s)
	}
}

func (s *ForCondStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitForCondStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	*StmtContext
	VALUE IExprContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ReturnStmtContext) GetVALUE() IExprContext { return s.VALUE }

func (s *ReturnStmtContext) SetVALUE(v IExprContext) { s.VALUE = v }

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RubaParserRULE_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewImportStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Match(RubaParserT__7)
		}
		{
			p.SetState(54)

			var _x = p.StringType()

			localctx.(*ImportStmtContext).NAME = _x
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RubaParserT__8 {
			{
				p.SetState(55)
				p.Match(RubaParserT__8)
			}
			{
				p.SetState(56)

				var _x = p.Ident()

				localctx.(*ImportStmtContext).ALIAS = _x
			}

		}

	case 2:
		localctx = NewDeclareStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RubaParserT__9 {
			{
				p.SetState(59)

				var _m = p.Match(RubaParserT__9)

				localctx.(*DeclareStmtContext).PUB = _m
			}

		}
		{
			p.SetState(62)

			var _x = p.Ident()

			localctx.(*DeclareStmtContext).NAME = _x
		}
		{
			p.SetState(63)
			p.Match(RubaParserT__10)
		}
		{
			p.SetState(64)

			var _x = p.expr(0)

			localctx.(*DeclareStmtContext).VALUE = _x
		}

	case 3:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)

			var _x = p.Ident()

			localctx.(*AssignStmtContext).ROOT = _x
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RubaParserT__4 || _la == RubaParserT__5 {
			{
				p.SetState(67)

				var _x = p.Nestable()

				localctx.(*AssignStmtContext)._nestable = _x
			}
			localctx.(*AssignStmtContext).CHAIN = append(localctx.(*AssignStmtContext).CHAIN, localctx.(*AssignStmtContext)._nestable)

			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(73)
			p.Match(RubaParserT__11)
		}
		{
			p.SetState(74)

			var _x = p.expr(0)

			localctx.(*AssignStmtContext).VALUE = _x
		}

	case 4:
		localctx = NewFnCallStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.FnCall()
		}

	case 5:
		localctx = NewFnDeclareStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(77)
			p.Match(RubaParserT__12)
		}
		{
			p.SetState(78)

			var _x = p.Ident()

			localctx.(*FnDeclareStmtContext).NAME = _x
		}
		{
			p.SetState(79)
			p.Match(RubaParserT__13)
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RubaParserT__15 || _la == RubaParserIdent {
			p.SetState(95)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RubaParserIdent:
				{
					p.SetState(80)

					var _x = p.Ident()

					localctx.(*FnDeclareStmtContext)._ident = _x
				}
				localctx.(*FnDeclareStmtContext).ARGS = append(localctx.(*FnDeclareStmtContext).ARGS, localctx.(*FnDeclareStmtContext)._ident)
				p.SetState(85)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(81)
							p.Match(RubaParserT__14)
						}
						{
							p.SetState(82)

							var _x = p.Ident()

							localctx.(*FnDeclareStmtContext)._ident = _x
						}
						localctx.(*FnDeclareStmtContext).ARGS = append(localctx.(*FnDeclareStmtContext).ARGS, localctx.(*FnDeclareStmtContext)._ident)

					}
					p.SetState(87)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
				}
				p.SetState(91)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == RubaParserT__14 {
					{
						p.SetState(88)
						p.Match(RubaParserT__14)
					}
					{
						p.SetState(89)
						p.Match(RubaParserT__15)
					}
					{
						p.SetState(90)

						var _x = p.Ident()

						localctx.(*FnDeclareStmtContext).RESTARG = _x
					}

				}

			case RubaParserT__15:
				{
					p.SetState(93)
					p.Match(RubaParserT__15)
				}
				{
					p.SetState(94)

					var _x = p.Ident()

					localctx.(*FnDeclareStmtContext).RESTARG = _x
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		{
			p.SetState(99)
			p.Match(RubaParserT__16)
		}
		{
			p.SetState(100)
			p.Match(RubaParserT__17)
		}
		{
			p.SetState(101)

			var _x = p.Block()

			localctx.(*FnDeclareStmtContext).BODY = _x
		}
		{
			p.SetState(102)
			p.Match(RubaParserT__18)
		}

	case 6:
		localctx = NewForInStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(104)
			p.Match(RubaParserT__19)
		}
		{
			p.SetState(105)

			var _x = p.Ident()

			localctx.(*ForInStmtContext)._ident = _x
		}
		localctx.(*ForInStmtContext).NAMES = append(localctx.(*ForInStmtContext).NAMES, localctx.(*ForInStmtContext)._ident)
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RubaParserT__14 {
			{
				p.SetState(106)
				p.Match(RubaParserT__14)
			}
			{
				p.SetState(107)

				var _x = p.Ident()

				localctx.(*ForInStmtContext)._ident = _x
			}
			localctx.(*ForInStmtContext).NAMES = append(localctx.(*ForInStmtContext).NAMES, localctx.(*ForInStmtContext)._ident)

			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(113)
			p.Match(RubaParserT__20)
		}
		{
			p.SetState(114)

			var _x = p.expr(0)

			localctx.(*ForInStmtContext).TARGET = _x
		}
		{
			p.SetState(115)
			p.Match(RubaParserT__17)
		}
		{
			p.SetState(116)

			var _x = p.Block()

			localctx.(*ForInStmtContext).BODY = _x
		}
		{
			p.SetState(117)
			p.Match(RubaParserT__18)
		}

	case 7:
		localctx = NewForCondStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(119)
			p.Match(RubaParserT__19)
		}
		{
			p.SetState(120)

			var _x = p.expr(0)

			localctx.(*ForCondStmtContext).COND = _x
		}
		{
			p.SetState(121)
			p.Match(RubaParserT__17)
		}
		{
			p.SetState(122)

			var _x = p.Block()

			localctx.(*ForCondStmtContext).BODY = _x
		}
		{
			p.SetState(123)
			p.Match(RubaParserT__18)
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RubaParserT__21 {
			{
				p.SetState(124)
				p.Match(RubaParserT__21)
			}
			{
				p.SetState(125)
				p.Match(RubaParserT__17)
			}
			{
				p.SetState(126)

				var _x = p.Block()

				localctx.(*ForCondStmtContext).ELSEBODY = _x
			}
			{
				p.SetState(127)
				p.Match(RubaParserT__18)
			}

		}

	case 8:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(131)
			p.Match(RubaParserT__22)
		}
		{
			p.SetState(132)

			var _x = p.expr(0)

			localctx.(*IfStmtContext).COND = _x
		}
		{
			p.SetState(133)
			p.Match(RubaParserT__17)
		}
		{
			p.SetState(134)

			var _x = p.Block()

			localctx.(*IfStmtContext).BODY = _x
		}
		{
			p.SetState(135)
			p.Match(RubaParserT__18)
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(136)
					p.Match(RubaParserT__21)
				}
				{
					p.SetState(137)
					p.Match(RubaParserT__22)
				}
				{
					p.SetState(138)

					var _x = p.expr(0)

					localctx.(*IfStmtContext)._expr = _x
				}
				localctx.(*IfStmtContext).ELIFCOND = append(localctx.(*IfStmtContext).ELIFCOND, localctx.(*IfStmtContext)._expr)
				{
					p.SetState(139)
					p.Match(RubaParserT__17)
				}
				{
					p.SetState(140)

					var _x = p.Block()

					localctx.(*IfStmtContext)._block = _x
				}
				localctx.(*IfStmtContext).ELIFBODY = append(localctx.(*IfStmtContext).ELIFBODY, localctx.(*IfStmtContext)._block)
				{
					p.SetState(141)
					p.Match(RubaParserT__18)
				}

			}
			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RubaParserT__21 {
			{
				p.SetState(148)
				p.Match(RubaParserT__21)
			}
			{
				p.SetState(149)
				p.Match(RubaParserT__17)
			}
			{
				p.SetState(150)

				var _x = p.Block()

				localctx.(*IfStmtContext).ELSEBODY = _x
			}
			{
				p.SetState(151)
				p.Match(RubaParserT__18)
			}

		}

	case 9:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(155)
			p.Match(RubaParserT__23)
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(156)

				var _x = p.expr(0)

				localctx.(*ReturnStmtContext).VALUE = _x
			}

		}

	case 10:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(159)
			p.Match(RubaParserT__24)
		}

	case 11:
		localctx = NewExprStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(160)

			var _x = p.expr(0)

			localctx.(*ExprStmtContext).E = _x
		}

	}

	return localctx
}

// IFnArgContext is an interface to support dynamic dispatch.
type IFnArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetARG returns the ARG rule contexts.
	GetARG() IExprContext

	// GetSPREADABLE returns the SPREADABLE rule contexts.
	GetSPREADABLE() IExprContext

	// SetARG sets the ARG rule contexts.
	SetARG(IExprContext)

	// SetSPREADABLE sets the SPREADABLE rule contexts.
	SetSPREADABLE(IExprContext)

	// IsFnArgContext differentiates from other interfaces.
	IsFnArgContext()
}

type FnArgContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ARG        IExprContext
	SPREADABLE IExprContext
}

func NewEmptyFnArgContext() *FnArgContext {
	var p = new(FnArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_fnArg
	return p
}

func (*FnArgContext) IsFnArgContext() {}

func NewFnArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnArgContext {
	var p = new(FnArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_fnArg

	return p
}

func (s *FnArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FnArgContext) GetARG() IExprContext { return s.ARG }

func (s *FnArgContext) GetSPREADABLE() IExprContext { return s.SPREADABLE }

func (s *FnArgContext) SetARG(v IExprContext) { s.ARG = v }

func (s *FnArgContext) SetSPREADABLE(v IExprContext) { s.SPREADABLE = v }

func (s *FnArgContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FnArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterFnArg(s)
	}
}

func (s *FnArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitFnArg(s)
	}
}

func (s *FnArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitFnArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) FnArg() (localctx IFnArgContext) {
	localctx = NewFnArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RubaParserRULE_fnArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)

			var _x = p.expr(0)

			localctx.(*FnArgContext).ARG = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)

			var _x = p.expr(0)

			localctx.(*FnArgContext).SPREADABLE = _x
		}
		{
			p.SetState(165)
			p.Match(RubaParserT__15)
		}

	}

	return localctx
}

// IFnCallContext is an interface to support dynamic dispatch.
type IFnCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNAME returns the NAME rule contexts.
	GetNAME() IIdentContext

	// GetEXPR returns the EXPR rule contexts.
	GetEXPR() IExprContext

	// Get_fnArg returns the _fnArg rule contexts.
	Get_fnArg() IFnArgContext

	// SetNAME sets the NAME rule contexts.
	SetNAME(IIdentContext)

	// SetEXPR sets the EXPR rule contexts.
	SetEXPR(IExprContext)

	// Set_fnArg sets the _fnArg rule contexts.
	Set_fnArg(IFnArgContext)

	// GetARGS returns the ARGS rule context list.
	GetARGS() []IFnArgContext

	// SetARGS sets the ARGS rule context list.
	SetARGS([]IFnArgContext)

	// IsFnCallContext differentiates from other interfaces.
	IsFnCallContext()
}

type FnCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	NAME   IIdentContext
	EXPR   IExprContext
	_fnArg IFnArgContext
	ARGS   []IFnArgContext
}

func NewEmptyFnCallContext() *FnCallContext {
	var p = new(FnCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_fnCall
	return p
}

func (*FnCallContext) IsFnCallContext() {}

func NewFnCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnCallContext {
	var p = new(FnCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_fnCall

	return p
}

func (s *FnCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FnCallContext) GetNAME() IIdentContext { return s.NAME }

func (s *FnCallContext) GetEXPR() IExprContext { return s.EXPR }

func (s *FnCallContext) Get_fnArg() IFnArgContext { return s._fnArg }

func (s *FnCallContext) SetNAME(v IIdentContext) { s.NAME = v }

func (s *FnCallContext) SetEXPR(v IExprContext) { s.EXPR = v }

func (s *FnCallContext) Set_fnArg(v IFnArgContext) { s._fnArg = v }

func (s *FnCallContext) GetARGS() []IFnArgContext { return s.ARGS }

func (s *FnCallContext) SetARGS(v []IFnArgContext) { s.ARGS = v }

func (s *FnCallContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FnCallContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FnCallContext) AllFnArg() []IFnArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFnArgContext)(nil)).Elem())
	var tst = make([]IFnArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFnArgContext)
		}
	}

	return tst
}

func (s *FnCallContext) FnArg(i int) IFnArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFnArgContext)
}

func (s *FnCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterFnCall(s)
	}
}

func (s *FnCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitFnCall(s)
	}
}

func (s *FnCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitFnCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) FnCall() (localctx IFnCallContext) {
	localctx = NewFnCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RubaParserRULE_fnCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(169)

			var _x = p.Ident()

			localctx.(*FnCallContext).NAME = _x
		}

	case 2:
		{
			p.SetState(170)

			var _x = p.expr(0)

			localctx.(*FnCallContext).EXPR = _x
		}

	}
	{
		p.SetState(173)
		p.Match(RubaParserT__13)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RubaParserT__1)|(1<<RubaParserT__5)|(1<<RubaParserT__13)|(1<<RubaParserT__17)|(1<<RubaParserT__27)|(1<<RubaParserT__28)|(1<<RubaParserT__29)|(1<<RubaParserT__30))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(RubaParserIdent-41))|(1<<(RubaParserInt-41))|(1<<(RubaParserStr-41))|(1<<(RubaParserFloat-41)))) != 0) {
		{
			p.SetState(174)

			var _x = p.FnArg()

			localctx.(*FnCallContext)._fnArg = _x
		}
		localctx.(*FnCallContext).ARGS = append(localctx.(*FnCallContext).ARGS, localctx.(*FnCallContext)._fnArg)
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RubaParserT__14 {
			{
				p.SetState(175)
				p.Match(RubaParserT__14)
			}
			{
				p.SetState(176)

				var _x = p.FnArg()

				localctx.(*FnCallContext)._fnArg = _x
			}
			localctx.(*FnCallContext).ARGS = append(localctx.(*FnCallContext).ARGS, localctx.(*FnCallContext)._fnArg)

			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(184)
		p.Match(RubaParserT__16)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RubaParserT__25 {
		{
			p.SetState(185)
			p.Match(RubaParserT__25)
		}

	}

	return localctx
}

// IObjectFieldContext is an interface to support dynamic dispatch.
type IObjectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSTATIC returns the STATIC rule contexts.
	GetSTATIC() IIdentContext

	// GetDYNAMIC returns the DYNAMIC rule contexts.
	GetDYNAMIC() IExprContext

	// GetVALUE returns the VALUE rule contexts.
	GetVALUE() IExprContext

	// SetSTATIC sets the STATIC rule contexts.
	SetSTATIC(IIdentContext)

	// SetDYNAMIC sets the DYNAMIC rule contexts.
	SetDYNAMIC(IExprContext)

	// SetVALUE sets the VALUE rule contexts.
	SetVALUE(IExprContext)

	// IsObjectFieldContext differentiates from other interfaces.
	IsObjectFieldContext()
}

type ObjectFieldContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	STATIC  IIdentContext
	DYNAMIC IExprContext
	VALUE   IExprContext
}

func NewEmptyObjectFieldContext() *ObjectFieldContext {
	var p = new(ObjectFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_objectField
	return p
}

func (*ObjectFieldContext) IsObjectFieldContext() {}

func NewObjectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectFieldContext {
	var p = new(ObjectFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_objectField

	return p
}

func (s *ObjectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectFieldContext) GetSTATIC() IIdentContext { return s.STATIC }

func (s *ObjectFieldContext) GetDYNAMIC() IExprContext { return s.DYNAMIC }

func (s *ObjectFieldContext) GetVALUE() IExprContext { return s.VALUE }

func (s *ObjectFieldContext) SetSTATIC(v IIdentContext) { s.STATIC = v }

func (s *ObjectFieldContext) SetDYNAMIC(v IExprContext) { s.DYNAMIC = v }

func (s *ObjectFieldContext) SetVALUE(v IExprContext) { s.VALUE = v }

func (s *ObjectFieldContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ObjectFieldContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ObjectFieldContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ObjectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterObjectField(s)
	}
}

func (s *ObjectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitObjectField(s)
	}
}

func (s *ObjectFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitObjectField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) ObjectField() (localctx IObjectFieldContext) {
	localctx = NewObjectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RubaParserRULE_objectField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RubaParserIdent:
		{
			p.SetState(188)

			var _x = p.Ident()

			localctx.(*ObjectFieldContext).STATIC = _x
		}

	case RubaParserT__5:
		{
			p.SetState(189)
			p.Match(RubaParserT__5)
		}
		{
			p.SetState(190)

			var _x = p.expr(0)

			localctx.(*ObjectFieldContext).DYNAMIC = _x
		}
		{
			p.SetState(191)
			p.Match(RubaParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(195)
		p.Match(RubaParserT__26)
	}
	{
		p.SetState(196)

		var _x = p.expr(0)

		localctx.(*ObjectFieldContext).VALUE = _x
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RubaParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RubaParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolExprContext struct {
	*ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringExprContext struct {
	*ExprContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) StringType() IStringTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringTypeContext)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualExprContext struct {
	*ExprContext
	L IExprContext
	R IExprContext
}

func NewEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualExprContext {
	var p = new(EqualExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualExprContext) GetL() IExprContext { return s.L }

func (s *EqualExprContext) GetR() IExprContext { return s.R }

func (s *EqualExprContext) SetL(v IExprContext) { s.L = v }

func (s *EqualExprContext) SetR(v IExprContext) { s.R = v }

func (s *EqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterEqualExpr(s)
	}
}

func (s *EqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitEqualExpr(s)
	}
}

func (s *EqualExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitEqualExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatExprContext struct {
	*ExprContext
}

func NewFloatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatExprContext {
	var p = new(FloatExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FloatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatExprContext) FloatType() IFloatTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatTypeContext)
}

func (s *FloatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterFloatExpr(s)
	}
}

func (s *FloatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitFloatExpr(s)
	}
}

func (s *FloatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitFloatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GtExprContext struct {
	*ExprContext
	L IExprContext
	R IExprContext
}

func NewGtExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GtExprContext {
	var p = new(GtExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GtExprContext) GetL() IExprContext { return s.L }

func (s *GtExprContext) GetR() IExprContext { return s.R }

func (s *GtExprContext) SetL(v IExprContext) { s.L = v }

func (s *GtExprContext) SetR(v IExprContext) { s.R = v }

func (s *GtExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GtExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GtExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GtExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterGtExpr(s)
	}
}

func (s *GtExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitGtExpr(s)
	}
}

func (s *GtExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitGtExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LtExprContext struct {
	*ExprContext
	L IExprContext
	R IExprContext
}

func NewLtExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LtExprContext {
	var p = new(LtExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LtExprContext) GetL() IExprContext { return s.L }

func (s *LtExprContext) GetR() IExprContext { return s.R }

func (s *LtExprContext) SetL(v IExprContext) { s.L = v }

func (s *LtExprContext) SetR(v IExprContext) { s.R = v }

func (s *LtExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LtExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LtExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LtExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterLtExpr(s)
	}
}

func (s *LtExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitLtExpr(s)
	}
}

func (s *LtExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitLtExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexExprContext struct {
	*ExprContext
	PARENT IExprContext
	CHILD  IExprContext
}

func NewIndexExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExprContext {
	var p = new(IndexExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IndexExprContext) GetPARENT() IExprContext { return s.PARENT }

func (s *IndexExprContext) GetCHILD() IExprContext { return s.CHILD }

func (s *IndexExprContext) SetPARENT(v IExprContext) { s.PARENT = v }

func (s *IndexExprContext) SetCHILD(v IExprContext) { s.CHILD = v }

func (s *IndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IndexExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterIndexExpr(s)
	}
}

func (s *IndexExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitIndexExpr(s)
	}
}

func (s *IndexExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitIndexExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullExprContext struct {
	*ExprContext
}

func NewNullExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExprContext {
	var p = new(NullExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NullExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterNullExpr(s)
	}
}

func (s *NullExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitNullExpr(s)
	}
}

func (s *NullExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitNullExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ObjectExprContext struct {
	*ExprContext
	_objectField IObjectFieldContext
	ELMS         []IObjectFieldContext
}

func NewObjectExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectExprContext {
	var p = new(ObjectExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ObjectExprContext) Get_objectField() IObjectFieldContext { return s._objectField }

func (s *ObjectExprContext) Set_objectField(v IObjectFieldContext) { s._objectField = v }

func (s *ObjectExprContext) GetELMS() []IObjectFieldContext { return s.ELMS }

func (s *ObjectExprContext) SetELMS(v []IObjectFieldContext) { s.ELMS = v }

func (s *ObjectExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectExprContext) AllObjectField() []IObjectFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem())
	var tst = make([]IObjectFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectFieldContext)
		}
	}

	return tst
}

func (s *ObjectExprContext) ObjectField(i int) IObjectFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectFieldContext)
}

func (s *ObjectExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterObjectExpr(s)
	}
}

func (s *ObjectExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitObjectExpr(s)
	}
}

func (s *ObjectExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitObjectExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExprContext struct {
	*ExprContext
	NAME   IExprContext
	_fnArg IFnArgContext
	ARGS   []IFnArgContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CallExprContext) GetNAME() IExprContext { return s.NAME }

func (s *CallExprContext) Get_fnArg() IFnArgContext { return s._fnArg }

func (s *CallExprContext) SetNAME(v IExprContext) { s.NAME = v }

func (s *CallExprContext) Set_fnArg(v IFnArgContext) { s._fnArg = v }

func (s *CallExprContext) GetARGS() []IFnArgContext { return s.ARGS }

func (s *CallExprContext) SetARGS(v []IFnArgContext) { s.ARGS = v }

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallExprContext) AllFnArg() []IFnArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFnArgContext)(nil)).Elem())
	var tst = make([]IFnArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFnArgContext)
		}
	}

	return tst
}

func (s *CallExprContext) FnArg(i int) IFnArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFnArgContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	*ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryExprContext struct {
	*ExprContext
	COND  IExprContext
	TRUE  IExprContext
	FALSE IExprContext
}

func NewTernaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExprContext {
	var p = new(TernaryExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TernaryExprContext) GetCOND() IExprContext { return s.COND }

func (s *TernaryExprContext) GetTRUE() IExprContext { return s.TRUE }

func (s *TernaryExprContext) GetFALSE() IExprContext { return s.FALSE }

func (s *TernaryExprContext) SetCOND(v IExprContext) { s.COND = v }

func (s *TernaryExprContext) SetTRUE(v IExprContext) { s.TRUE = v }

func (s *TernaryExprContext) SetFALSE(v IExprContext) { s.FALSE = v }

func (s *TernaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TernaryExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TernaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterTernaryExpr(s)
	}
}

func (s *TernaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitTernaryExpr(s)
	}
}

func (s *TernaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitTernaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentExprContext struct {
	*ExprContext
}

func NewIdentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentExprContext {
	var p = new(IdentExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IdentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentExprContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *IdentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterIdentExpr(s)
	}
}

func (s *IdentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitIdentExpr(s)
	}
}

func (s *IdentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitIdentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecExprContext struct {
	*ExprContext
}

func NewDecExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecExprContext {
	var p = new(DecExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DecExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecExprContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *DecExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterDecExpr(s)
	}
}

func (s *DecExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitDecExpr(s)
	}
}

func (s *DecExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitDecExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncExprContext struct {
	*ExprContext
	L IExprContext
}

func NewIncExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncExprContext {
	var p = new(IncExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IncExprContext) GetL() IExprContext { return s.L }

func (s *IncExprContext) SetL(v IExprContext) { s.L = v }

func (s *IncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IncExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterIncExpr(s)
	}
}

func (s *IncExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitIncExpr(s)
	}
}

func (s *IncExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitIncExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddOrSubExprContext struct {
	*ExprContext
	L IExprContext
	S antlr.Token
	R IExprContext
}

func NewAddOrSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddOrSubExprContext {
	var p = new(AddOrSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddOrSubExprContext) GetS() antlr.Token { return s.S }

func (s *AddOrSubExprContext) SetS(v antlr.Token) { s.S = v }

func (s *AddOrSubExprContext) GetL() IExprContext { return s.L }

func (s *AddOrSubExprContext) GetR() IExprContext { return s.R }

func (s *AddOrSubExprContext) SetL(v IExprContext) { s.L = v }

func (s *AddOrSubExprContext) SetR(v IExprContext) { s.R = v }

func (s *AddOrSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOrSubExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddOrSubExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddOrSubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterAddOrSubExpr(s)
	}
}

func (s *AddOrSubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitAddOrSubExpr(s)
	}
}

func (s *AddOrSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitAddOrSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DotExprContext struct {
	*ExprContext
	PARENT IExprContext
	CHILD  IIdentContext
}

func NewDotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotExprContext {
	var p = new(DotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DotExprContext) GetPARENT() IExprContext { return s.PARENT }

func (s *DotExprContext) GetCHILD() IIdentContext { return s.CHILD }

func (s *DotExprContext) SetPARENT(v IExprContext) { s.PARENT = v }

func (s *DotExprContext) SetCHILD(v IIdentContext) { s.CHILD = v }

func (s *DotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DotExprContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *DotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterDotExpr(s)
	}
}

func (s *DotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitDotExpr(s)
	}
}

func (s *DotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitDotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotEqualExprContext struct {
	*ExprContext
	L IExprContext
	R IExprContext
}

func NewNotEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualExprContext {
	var p = new(NotEqualExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotEqualExprContext) GetL() IExprContext { return s.L }

func (s *NotEqualExprContext) GetR() IExprContext { return s.R }

func (s *NotEqualExprContext) SetL(v IExprContext) { s.L = v }

func (s *NotEqualExprContext) SetR(v IExprContext) { s.R = v }

func (s *NotEqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *NotEqualExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotEqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterNotEqualExpr(s)
	}
}

func (s *NotEqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitNotEqualExpr(s)
	}
}

func (s *NotEqualExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitNotEqualExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LteExprContext struct {
	*ExprContext
	L IExprContext
	R IExprContext
}

func NewLteExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LteExprContext {
	var p = new(LteExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LteExprContext) GetL() IExprContext { return s.L }

func (s *LteExprContext) GetR() IExprContext { return s.R }

func (s *LteExprContext) SetL(v IExprContext) { s.L = v }

func (s *LteExprContext) SetR(v IExprContext) { s.R = v }

func (s *LteExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LteExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LteExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LteExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterLteExpr(s)
	}
}

func (s *LteExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitLteExpr(s)
	}
}

func (s *LteExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitLteExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GteExprContext struct {
	*ExprContext
	L IExprContext
	R IExprContext
}

func NewGteExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GteExprContext {
	var p = new(GteExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GteExprContext) GetL() IExprContext { return s.L }

func (s *GteExprContext) GetR() IExprContext { return s.R }

func (s *GteExprContext) SetL(v IExprContext) { s.L = v }

func (s *GteExprContext) SetR(v IExprContext) { s.R = v }

func (s *GteExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GteExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GteExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GteExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterGteExpr(s)
	}
}

func (s *GteExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitGteExpr(s)
	}
}

func (s *GteExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitGteExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupExprContext struct {
	*ExprContext
	EXPR IExprContext
}

func NewGroupExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupExprContext {
	var p = new(GroupExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GroupExprContext) GetEXPR() IExprContext { return s.EXPR }

func (s *GroupExprContext) SetEXPR(v IExprContext) { s.EXPR = v }

func (s *GroupExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GroupExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterGroupExpr(s)
	}
}

func (s *GroupExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitGroupExpr(s)
	}
}

func (s *GroupExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitGroupExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegativeExprContext struct {
	*ExprContext
}

func NewNegativeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegativeExprContext {
	var p = new(NegativeExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NegativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegativeExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NegativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterNegativeExpr(s)
	}
}

func (s *NegativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitNegativeExpr(s)
	}
}

func (s *NegativeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitNegativeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayExprContext struct {
	*ExprContext
	_expr IExprContext
	ELMS  []IExprContext
}

func NewArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExprContext {
	var p = new(ArrayExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ArrayExprContext) Get_expr() IExprContext { return s._expr }

func (s *ArrayExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ArrayExprContext) GetELMS() []IExprContext { return s.ELMS }

func (s *ArrayExprContext) SetELMS(v []IExprContext) { s.ELMS = v }

func (s *ArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArrayExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterArrayExpr(s)
	}
}

func (s *ArrayExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitArrayExpr(s)
	}
}

func (s *ArrayExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitArrayExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulOrDivExprContext struct {
	*ExprContext
	L IExprContext
	S antlr.Token
	R IExprContext
}

func NewMulOrDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulOrDivExprContext {
	var p = new(MulOrDivExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulOrDivExprContext) GetS() antlr.Token { return s.S }

func (s *MulOrDivExprContext) SetS(v antlr.Token) { s.S = v }

func (s *MulOrDivExprContext) GetL() IExprContext { return s.L }

func (s *MulOrDivExprContext) GetR() IExprContext { return s.R }

func (s *MulOrDivExprContext) SetL(v IExprContext) { s.L = v }

func (s *MulOrDivExprContext) SetR(v IExprContext) { s.R = v }

func (s *MulOrDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOrDivExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulOrDivExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulOrDivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterMulOrDivExpr(s)
	}
}

func (s *MulOrDivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitMulOrDivExpr(s)
	}
}

func (s *MulOrDivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitMulOrDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntExprContext struct {
	*ExprContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) IntType() IIntTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntTypeContext)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RubaListener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RubaVisitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RubaParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *RubaParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, RubaParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGroupExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(199)
			p.Match(RubaParserT__13)
		}
		{
			p.SetState(200)

			var _x = p.expr(0)

			localctx.(*GroupExprContext).EXPR = _x
		}
		{
			p.SetState(201)
			p.Match(RubaParserT__16)
		}

	case 2:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(203)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RubaParserT__27 || _la == RubaParserT__28) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 3:
		localctx = NewNullExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(204)
			p.Match(RubaParserT__29)
		}

	case 4:
		localctx = NewIdentExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(205)
			p.Ident()
		}

	case 5:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(206)
			p.IntType()
		}

	case 6:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(207)
			p.FloatType()
		}

	case 7:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(208)
			p.StringType()
		}

	case 8:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(209)
			p.Match(RubaParserT__30)
		}
		{
			p.SetState(210)
			p.expr(15)
		}

	case 9:
		localctx = NewNegativeExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(211)
			p.Match(RubaParserT__1)
		}
		{
			p.SetState(212)
			p.expr(14)
		}

	case 10:
		localctx = NewDecExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(213)
			p.Ident()
		}
		{
			p.SetState(214)
			p.Match(RubaParserT__38)
		}

	case 11:
		localctx = NewArrayExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(216)
			p.Match(RubaParserT__5)
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RubaParserT__1)|(1<<RubaParserT__5)|(1<<RubaParserT__13)|(1<<RubaParserT__17)|(1<<RubaParserT__27)|(1<<RubaParserT__28)|(1<<RubaParserT__29)|(1<<RubaParserT__30))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(RubaParserIdent-41))|(1<<(RubaParserInt-41))|(1<<(RubaParserStr-41))|(1<<(RubaParserFloat-41)))) != 0) {
			{
				p.SetState(217)

				var _x = p.expr(0)

				localctx.(*ArrayExprContext)._expr = _x
			}
			localctx.(*ArrayExprContext).ELMS = append(localctx.(*ArrayExprContext).ELMS, localctx.(*ArrayExprContext)._expr)
			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == RubaParserT__14 {
				{
					p.SetState(218)
					p.Match(RubaParserT__14)
				}
				{
					p.SetState(219)

					var _x = p.expr(0)

					localctx.(*ArrayExprContext)._expr = _x
				}
				localctx.(*ArrayExprContext).ELMS = append(localctx.(*ArrayExprContext).ELMS, localctx.(*ArrayExprContext)._expr)

				p.SetState(224)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(227)
			p.Match(RubaParserT__6)
		}

	case 12:
		localctx = NewObjectExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(228)
			p.Match(RubaParserT__17)
		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RubaParserT__5 || _la == RubaParserIdent {
			{
				p.SetState(229)

				var _x = p.ObjectField()

				localctx.(*ObjectExprContext)._objectField = _x
			}
			localctx.(*ObjectExprContext).ELMS = append(localctx.(*ObjectExprContext).ELMS, localctx.(*ObjectExprContext)._objectField)
			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(230)
						p.Match(RubaParserT__14)
					}
					{
						p.SetState(231)

						var _x = p.ObjectField()

						localctx.(*ObjectExprContext)._objectField = _x
					}
					localctx.(*ObjectExprContext).ELMS = append(localctx.(*ObjectExprContext).ELMS, localctx.(*ObjectExprContext)._objectField)

				}
				p.SetState(236)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
			}
			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == RubaParserT__14 {
				{
					p.SetState(237)
					p.Match(RubaParserT__14)
				}

			}

		}
		{
			p.SetState(242)
			p.Match(RubaParserT__18)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(298)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulOrDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulOrDivExprContext).L = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				p.SetState(246)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*MulOrDivExprContext).S = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == RubaParserT__2 || _la == RubaParserT__3) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*MulOrDivExprContext).S = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(247)

					var _x = p.expr(14)

					localctx.(*MulOrDivExprContext).R = _x
				}

			case 2:
				localctx = NewAddOrSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddOrSubExprContext).L = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(249)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AddOrSubExprContext).S = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == RubaParserT__0 || _la == RubaParserT__1) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AddOrSubExprContext).S = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(250)

					var _x = p.expr(13)

					localctx.(*AddOrSubExprContext).R = _x
				}

			case 3:
				localctx = NewEqualExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqualExprContext).L = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(252)
					p.Match(RubaParserT__31)
				}
				{
					p.SetState(253)

					var _x = p.expr(12)

					localctx.(*EqualExprContext).R = _x
				}

			case 4:
				localctx = NewNotEqualExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*NotEqualExprContext).L = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(255)
					p.Match(RubaParserT__32)
				}
				{
					p.SetState(256)

					var _x = p.expr(11)

					localctx.(*NotEqualExprContext).R = _x
				}

			case 5:
				localctx = NewGtExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GtExprContext).L = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(258)
					p.Match(RubaParserT__33)
				}
				{
					p.SetState(259)

					var _x = p.expr(10)

					localctx.(*GtExprContext).R = _x
				}

			case 6:
				localctx = NewLteExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LteExprContext).L = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(261)
					p.Match(RubaParserT__34)
				}
				{
					p.SetState(262)

					var _x = p.expr(9)

					localctx.(*LteExprContext).R = _x
				}

			case 7:
				localctx = NewLtExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LtExprContext).L = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(264)
					p.Match(RubaParserT__35)
				}
				{
					p.SetState(265)

					var _x = p.expr(8)

					localctx.(*LtExprContext).R = _x
				}

			case 8:
				localctx = NewGteExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GteExprContext).L = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(267)
					p.Match(RubaParserT__36)
				}
				{
					p.SetState(268)

					var _x = p.expr(7)

					localctx.(*GteExprContext).R = _x
				}

			case 9:
				localctx = NewTernaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TernaryExprContext).COND = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(270)
					p.Match(RubaParserT__39)
				}
				{
					p.SetState(271)

					var _x = p.expr(0)

					localctx.(*TernaryExprContext).TRUE = _x
				}
				{
					p.SetState(272)
					p.Match(RubaParserT__26)
				}
				{
					p.SetState(273)

					var _x = p.expr(2)

					localctx.(*TernaryExprContext).FALSE = _x
				}

			case 10:
				localctx = NewCallExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CallExprContext).NAME = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(276)
					p.Match(RubaParserT__13)
				}
				p.SetState(285)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RubaParserT__1)|(1<<RubaParserT__5)|(1<<RubaParserT__13)|(1<<RubaParserT__17)|(1<<RubaParserT__27)|(1<<RubaParserT__28)|(1<<RubaParserT__29)|(1<<RubaParserT__30))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(RubaParserIdent-41))|(1<<(RubaParserInt-41))|(1<<(RubaParserStr-41))|(1<<(RubaParserFloat-41)))) != 0) {
					{
						p.SetState(277)

						var _x = p.FnArg()

						localctx.(*CallExprContext)._fnArg = _x
					}
					localctx.(*CallExprContext).ARGS = append(localctx.(*CallExprContext).ARGS, localctx.(*CallExprContext)._fnArg)
					p.SetState(282)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == RubaParserT__14 {
						{
							p.SetState(278)
							p.Match(RubaParserT__14)
						}
						{
							p.SetState(279)

							var _x = p.FnArg()

							localctx.(*CallExprContext)._fnArg = _x
						}
						localctx.(*CallExprContext).ARGS = append(localctx.(*CallExprContext).ARGS, localctx.(*CallExprContext)._fnArg)

						p.SetState(284)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(287)
					p.Match(RubaParserT__16)
				}

			case 11:
				localctx = NewIndexExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*IndexExprContext).PARENT = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(289)
					p.Match(RubaParserT__5)
				}
				{
					p.SetState(290)

					var _x = p.expr(0)

					localctx.(*IndexExprContext).CHILD = _x
				}
				{
					p.SetState(291)
					p.Match(RubaParserT__6)
				}

			case 12:
				localctx = NewDotExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DotExprContext).PARENT = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(294)
					p.Match(RubaParserT__4)
				}
				{
					p.SetState(295)

					var _x = p.Ident()

					localctx.(*DotExprContext).CHILD = _x
				}

			case 13:
				localctx = NewIncExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*IncExprContext).L = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RubaParserRULE_expr)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(297)
					p.Match(RubaParserT__37)
				}

			}

		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

func (p *RubaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RubaParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
