import React from 'react';
import { Form, Input, Button, Card, Checkbox } from 'antd';
import { authenticationService } from '../services/authentication';
import { route } from '../config';
import { useHistory } from 'react-router';

function Register() {
  let history = useHistory();
  const onFinish = async (values: any) => {
    const matchPass = values.password === values.confirmPassword;
    if (matchPass) {
      const success = await authenticationService.register(values.email, values.password);
      if (success) {
        history.push(route.login);
      }
      //TODO alert not success return
    }
    //TODO alert not match
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

          <Form.Item
            label="ยืนยันรหัสผ่าน"
            name="confirmPassword"
            rules={[{ required: true, message: 'Please input your confirm password!' }]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item
            name="termAndCon"
            valuePropName="checked"
            rules={[{ required: true, message: 'term and condition is require' }]}
          >
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
  );
}

export default Register;
