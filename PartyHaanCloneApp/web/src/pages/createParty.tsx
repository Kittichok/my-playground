import React from 'react';
import { Form, Input, Button, Card, Layout, InputNumber } from 'antd';
import * as partyServices from '../services/party';
import { useHistory } from 'react-router-dom';
import { route } from '../config';

const { Header, Content } = Layout;

function CreateParty() {
  let history = useHistory();
  const onFinish = async (values: any) => {
    const success = await partyServices.create(values.partyName, values.number);
    if (success) {
      history.push(route.listing);
    }
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
    //TODO refactor layout to private route
    <Layout>
      <Header style={{ backgroundColor: 'rgb(24, 144, 255)', color: 'white' }}>สร้างปาร์ตี้</Header>
      <Content style={{ backgroundColor: 'whitesmoke' }}>
        <Card>
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
              rules={[
                {
                  required: true,
                  message: 'Please input your number of people!',
                },
              ]}
            >
              <InputNumber style={{ width: '100%' }} />
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
  );
}

export default CreateParty;
