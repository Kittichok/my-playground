import React from 'react';
import logo from '../assets/logo.svg';
import { Form, Input, Button, Card } from 'antd';
import { useHistory } from 'react-router-dom';
import { authenticationService } from '../services/authentication';
import { route } from '../config';

function Login() {
  let history = useHistory();
  const onFinish = async (values: any) => {
    const token = await authenticationService.login(values.email, values.password);
    if (token) {
      history.push(route.listing);
    } else {
      //TODO popup login fail
      console.log('login fail');
    }
  };

  const onRegister = () => {
    history.push(route.register);
  };

  return (
    <div>
      <Card>
        <img src={logo} className="App-logo" alt="logo" />
        <Form layout="vertical" initialValues={{ remember: true }} onFinish={onFinish}>
          <Form.Item
            label="อีเมล์"
            name="email"
            rules={[
              {
                required: true,
                message: 'Please input your email!',
                type: 'email',
              },
            ]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            label="รหัสผ่าน"
            name="password"
            rules={[{ required: true, message: 'Please input your password!' }]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit">
              เข้าสู่ระบบ
            </Button>
          </Form.Item>

          <Form.Item>
            <Button type="primary" onClick={onRegister}>
              สร้างบัญชีผู้ใช้
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
}

export default Login;
