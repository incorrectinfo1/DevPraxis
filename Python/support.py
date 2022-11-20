import base64

#snagged from dotpeak and userinfo.exe
enc_password = "0Nv32PTwgYjzg9/8j5TbmvPd3e7WhtWWyuPsyO76/Y+U193E"
# private static byte[] key = Encoding.ASCII.GetBytes("armando");
key = b'armando'

# byte[] array = Convert.FromBase64String(Protected.enc_password);
# byte[] array2 = array;
array = base64.b64decode(enc_password)

# for (int i = 0; i < array.Length; i++){
#     array2[i] = (array[i] ^ Protected.key[i % Protected.key.Length] ^ 223);
# }
array2 = ''
for i in range(len(array)):
    array2 += chr(array[i] ^ key[i%len(key)] ^ 223)

print("Decoding: " + enc_password)
# return Encoding.Default.GetString(array2);
print("LDAP Key: " + array2)