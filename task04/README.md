switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough

cannot fallthrough in type switch  在switch类型判断模式中，不能使用fallthrough

select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测