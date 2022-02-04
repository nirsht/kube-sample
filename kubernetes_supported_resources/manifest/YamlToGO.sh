cd kubernetes_supported_resources/manifest
for file in ./*; do
    input="$(basename "$file")"
    extension="${input##*.}"
    filename="${input%.*}"
    if [ $extension = "yaml" ]
    then
        rm -f ./$filename.go
        echo "package manifest" >> ./$filename.go
        echo "" >> ./$filename.go
        echo "const $filename"' = `' >> ./$filename.go
        while IFS= read -r line
        do
            echo "$line" >> ./$filename.go
        done < "$input"
        echo '`' >> ./$filename.go
    fi
done

