import React from 'react';
import logo from '../assets/logo.svg';
import { Form, Input, Button, Card } from 'antd';

function Login() {
  const onFinish = (values: any) => {
    console.log('Success:', values);
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
    <div>
      <img src={logo} className="App-logo" alt="logo" />
      <Card style={{ width: 400 }}>
        <Form
          layout="vertical"
          initialValues={{ remember: true }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
        >
          <Form.Item
            label="อีเมล์"
            name="email"
            rules={[{ required: true, message: 'Please input your username!' }]}
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
            <Button type="primary">
              สร้างบัญชีผู้ใช้
            </Button>
          </Form.Item>

        </Form>
      </Card>
    </div>
  )
}

export default Login;