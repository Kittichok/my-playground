import crypto from 'crypto';

const genRandomString = (length) => {
  return crypto
    .randomBytes(Math.ceil(length / 2))
    .toString('hex')
    .slice(0, length);
};

const sha512 = (password: string, salt: string) => {
  var hash = crypto.createHmac('sha512', salt);
  hash.update(password);
  var value = hash.digest('hex');
  return value;
};

const saltHashPassword = (userpassword: string) => {
  var salt = genRandomString(16);
  var hash = sha512(userpassword, salt);
  return {
    hash,
    salt,
  };
};

const validateHashPassword = (password: string, salt: string, passwordHash: string) => {
  const hash = sha512(password, salt);
  return hash === passwordHash;
};

export { saltHashPassword, sha512, validateHashPassword };
