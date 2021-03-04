import React from 'react';
import { Form, Input, Button, Card, Checkbox } from 'antd';

function Register() {
  const onFinish = (values: any) => {
    console.log('Success:', values);
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
    <div>
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

          <Form.Item
            label="ยืนยันรหัสผ่าน"
            name="confirmPassword"
            rules={[{ required: true, message: 'Please input your password!' }]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item name="termAndCon" valuePropName="checked">
            <Checkbox>ฉันยอมรับเงื่อนไขและข้อตกลงเกี่ยวกับการใช้งาน</Checkbox>
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit">
              ยืนยัน
            </Button>
          </Form.Item>
          </Form>
      </Card>
    </div>
  )
}

export default Register;