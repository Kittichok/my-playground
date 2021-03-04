import React from 'react';
import { Form, Input, Button, Card, Layout } from 'antd';

const { Header, Content } = Layout;

function CreateParty() {
  const onFinish = (values: any) => {
    console.log('Success:', values);
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
    <Layout>
      <Header style={{ backgroundColor: 'whitesmoke' }}>สร้างปาร์ตี้</Header>
      <Content>
        <Card style={{ width: 400 }}>
          <Form
            layout="vertical"
            initialValues={{ remember: true }}
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
          >
            <Form.Item
              label="ชื่อปาร์ตี้"
              name="partyName"
              rules={[{ required: true, message: 'Please input your party name!' }]}
            >
              <Input />
            </Form.Item>

            <Form.Item
              label="จำนวนคนที่ขาด"
              name="number"
              rules={[{ required: true, message: 'Please input your number of people!' }]}
            >
              <Input />
            </Form.Item>

            <Form.Item>
              <Button type="primary" htmlType="submit">
                สร้างปาร์ตี้
            </Button>
            </Form.Item>
          </Form>
        </Card>
      </Content>
    </Layout>
  )
}

export default CreateParty;
