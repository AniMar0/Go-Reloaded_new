package reload

func Single_Cote(content string) string {
    new_content := Split_Single_Cote(content)

    if len(new_content) == 0 {
        return content
    }

    str := ""
    for _, arg := range new_content {
        if arg[0] == '\'' && arg[len(arg)-1] == '\'' && arg != "'" {
            str += "'" + RemoveSpace(arg[1:len(arg)-1]) + "' "
        } else {
            str += arg + " "
        }
    }

    return str
}

func Split_Single_Cote(Data string) []string {
    Words := []string{}
    firstindex := 0
    lastindex := 0
    inQuote := false // مؤشر إذا كنا داخل علامات الاقتباس
    counter := 0

    for i, char := range Data {
        if char == '\'' {
            // التحقق مما إذا كانت علامة الاقتباس بين حرفين أبجديين
            if i > 0 && i < len(Data)-1 && isAlpha(rune(Data[i-1])) && isAlpha(rune(Data[i+1])) {
                continue // تخطي علامة الاقتباس لأنها جزء من الكلمة
            }

            if inQuote {
                // إغلاق الاقتباس
                inQuote = false
                lastindex = i + 1
                Words = append(Words, Data[firstindex:lastindex]) // إضافة النص بين علامتي الاقتباس
                counter++
            } else {
                // فتح الاقتباس
                inQuote = true
                firstindex = i
                if lastindex < firstindex {
                    Words = append(Words, Data[lastindex:firstindex]) // إضافة النص خارج الاقتباس
                }
            }
        }
    }

    // إذا كان الاقتباس مفتوحًا ولم يغلق
    if inQuote {
        Words = append(Words, Data[firstindex:])
    } else if len(Data) > lastindex {
        Words = append(Words, Data[lastindex:])
    }

    //println(counter) // طباعة عدد علامات الاقتباس المغلقة
    return Words
}