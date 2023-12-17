package deepcopy

import (
	"reflect"
	"time"
)

// 深拷贝

func Copy(src interface{}) interface{} {
	origianl := reflect.ValueOf(src)
	dest := reflect.New(origianl.Type()).Elem()
	copyRecursive(origianl, dest)
	return dest.Interface()
}

// 将 src 复制到 dest （这里dest和src的类型是一致的）
func copyRecursive(src, dest reflect.Value) {
	switch src.Kind() {

	case reflect.Ptr: // 说明src是一个指针，不能直接将指针复制给dest；如果直接复制，src和dest就指向了同一个地址（就不算是深拷贝）
		original := src.Elem() // 指针指向的类型
		if src.IsNil() || !original.IsValid() {
			return
		}
		// 所以这里构造了一个新地址
		destValue := reflect.New(original.Type())

		// 这里将 original的值复制给 v.Elem()
		copyRecursive(original, destValue.Elem())
		// 让dest指向这个新地址
		dest.Set(destValue)

	case reflect.Interface:
		if src.IsNil() {
			return
		}

		// interface 包含的类型
		original := src.Elem()
		// 构造该类型的变量
		destValue := reflect.New(original.Type()).Elem()
		copyRecursive(original, destValue)
		// 保存到 dest中
		dest.Set(destValue)
	case reflect.Slice:
		if src.IsNil() {
			return
		}
		//第一种写法： 构造切片
		dest.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		// 对切片中的每个元素进行复制
		for i := 0; i < src.Len(); i++ {
			copyRecursive(src.Index(i), dest.Index(i))
		}

		// 第二种写法：通过append追加
		// for i := 0; i < src.Len(); i++ {
		// 	tmp := reflect.New(src.Index(i).Type()).Elem()
		// 	copyRecursive(src.Index(i), tmp)
		// 	dest.Set(reflect.Append(dest, tmp))
		// }

	case reflect.Map:
		if src.IsNil() {
			return
		}
		dest.Set(reflect.MakeMap(src.Type()))
		for _, key := range src.MapKeys() {

			originValue := src.MapIndex(key)
			destValue := reflect.New(originValue.Type()).Elem()
			copyRecursive(originValue, destValue)
			// 第一种写法：这里的key可能是任何类型
			destKey := Copy(key.Interface())
			dest.SetMapIndex(reflect.ValueOf(destKey), destValue)

			// 第二种写法：
			// destKey := reflect.New(key.Type()).Elem()
			// copyRecursive(key, destKey)
			// dest.SetMapIndex(destKey, destValue)

		}
	case reflect.Struct:
		t, ok := src.Interface().(time.Time)
		if ok {
			dest.Set(reflect.ValueOf(t))
			return
		} else {

			for i := 0; i < src.NumField(); i++ {
				p := src.Type().Field(i)
				if !p.IsExported() {
					continue
				}
				copyRecursive(src.Field(i), dest.Field(i))
			}
		}
	default:
		dest.Set(src)
	}
}
