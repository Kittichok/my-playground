import { Layout, Button, Row, Col } from 'antd';
import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import PartyCard from '../components/partyCard';
import { route } from '../config';
import * as partyService from '../services/party';

const { Header, Content } = Layout;

function PartyList() {
  let history = useHistory();
  const [partyListData, setPartyListData] = useState<any[]>();

  useEffect(() => {
    getPartyList();
  }, []);

  const getPartyList = async () => {
    const list = await partyService.getList();
    setPartyListData(list);
  };

  const createParty = () => {
    history.push(route.createParty);
  };

  const joinAction = async (partyID: number) => {
    console.log('click');

    const success = await partyService.join(partyID);
    if (success) {
      getPartyList();
    }
  };

  return (
    //TODO refactor layout to private route
    <Layout>
      <Header style={{ backgroundColor: 'rgb(24, 144, 255)', color: 'white' }}>
        ปาร์ตี้ทั้งหมด
      </Header>
      <Content style={{ backgroundColor: 'whitesmoke' }}>
        <Button type="primary" onClick={createParty} style={{ marginTop: '10px' }}>
          สร้างปาร์ตี้
        </Button>
        <Row gutter={[16, 16]} style={{ margin: '10px' }}>
          {partyListData
            ? partyListData.map((data) => (
                <Col span={12} key={data.id}>
                  <PartyCard
                    key={data.id}
                    img={data.img}
                    name={data.name}
                    totalMember={data.totalMember}
                    currentMember={data.currentMember}
                    joinAction={() => joinAction(data.id)}
                  />
                </Col>
              ))
            : null}
        </Row>
      </Content>
    </Layout>
  );
}

export default PartyList;
