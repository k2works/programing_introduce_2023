package main

import (
	"contens/chapter02/section004/hello"
	"contens/chapter02/section007/_bool"
	"contens/chapter02/section007/number"
	_string "contens/chapter02/section007/string"
	_type "contens/chapter02/section008/type"
	"contens/chapter02/section009/convert"
	"contens/chapter02/section010/arithmetic"
	"contens/chapter02/section010/arithmetic_type"
	"contens/chapter02/section010/compare"
	"contens/chapter02/section010/inc_dec"
	"contens/chapter02/section010/logical"
	"contens/chapter02/section012/abbreviation"
	_break "contens/chapter02/section016/break"
	"contens/chapter02/section016/condition"
	continue_ "contens/chapter02/section016/continue"
	_for "contens/chapter02/section016/for"
	"contens/chapter02/section016/infinity"
	_range "contens/chapter02/section016/range"
	range_string "contens/chapter02/section016/range_strng"
	"contens/chapter02/section017/falthrough"
	_if "contens/chapter02/section017/if"
	"contens/chapter02/section017/swithc_true"
	"contens/chapter02/section017/swithch"
	"contens/chapter02/section018/pointer"
	"contens/chapter02/section018/pointer_value"
	"contens/chapter02/section019/zero"
	_goto "contens/chapter02/section020/goto"
	"contens/chapter02/section020/label_break"
	"contens/chapter02/section020/label_continue"
	"contens/chapter03/section021/function"
	"contens/chapter03/section021/function_literal"
	"contens/chapter03/section021/function_type"
	"contens/chapter03/section021/multi"
	"contens/chapter03/section021/multi_param"
	_return "contens/chapter03/section021/return"
	"contens/chapter03/section021/variadic"
	"contens/chapter03/section022/method"
	"contens/chapter03/section022/method_value"
	"contens/chapter03/section022/receiver_ptr"
	_defer "contens/chapter03/section023/defer"
	"contens/chapter03/section023/defer_close"
	_struct "contens/chapter04/section024/struct"
	"contens/chapter04/section025/embed"
	"contens/chapter04/section026/embed_literal"
	"contens/chapter04/section026/struct_init"
	_interface "contens/chapter05/section027/interface"
	"contens/chapter05/section029/assertion"
	"contens/chapter05/section029/assertion2"
	"contens/chapter05/section029/type_switch"
	"contens/chapter06/section030/index"
	_len "contens/chapter06/section030/len"
	"contens/chapter06/section031/literal"
)

func main() {
	//Section004
	result := hello.Greeting()
	println(result)

	//Section007
	_bool.Exec()
	number.Exec()
	_string.Exec()

	//Section008
	_type.Exec()

	//Section009
	convert.Exec()

	//Section010
	arithmetic.Exec()
	arithmetic_type.Exec()
	compare.Exec()
	inc_dec.Exec()
	logical.Exec()

	//Section012
	abbreviation.Exec()

	//Section016
	_break.Exec()
	condition.Exec()
	continue_.Exec()
	_for.Exec()
	infinity.Exec()
	_range.Exec()
	range_string.Exec()

	//Section017
	falthrough.Exec()
	_if.Exec()
	swithc_true.Exec()
	swithch.Exec()

	//Section018
	pointer.Exec()
	pointer_value.Exec()

	//Section019
	zero.Exec()

	//Section020
	label_break.Exec()
	label_continue.Exec()
	_goto.Exec()

	//Section021
	function.Exec()
	multi.Exec()
	multi_param.Exec()
	variadic.Exec()
	_return.Exec()
	function_literal.Exec()
	function_type.Exec()

	//Section022
	method.Exec()
	receiver_ptr.Exec()
	method_value.Exec()

	//Section023
	_defer.Exec()
	defer_close.Exec()

	//Section024
	_struct.Exec()

	//Section025
	embed.Exec()

	//Section026
	struct_init.Exec()
	embed_literal.Exec()

	//Section027
	_interface.Exec()

	//Section029
	assertion.Exec()
	assertion2.Exec()
	type_switch.Exec()

	//Section030
	_len.Exec()
	index.Exec()

	//Section031
	literal.Exec()
}
