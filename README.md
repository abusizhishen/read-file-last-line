# read-file-last-line
read-file-last-line

# step
1. open file and get file size
2. readAt() form file end,  
3. loop 
	- read file like pieces, size from 64byte to 128byte to 256byte ...  
    - put contents to bytes.Buffer 
    - record content length
    
	until find line break
4. reassemble contents, like from [e,d,bc,a] to [a,bc,d,e]
