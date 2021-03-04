import { Card, Layout, Button } from 'antd';
import React from 'react';
import { useHistory } from 'react-router-dom';
import PartyCard from '../components/partyCard';
import * as partyService from '../services/party';

const { Header, Content } = Layout;

function PartyList() {
  let history = useHistory();
  const partyListData = partyService.getList();

  const createParty = () => {
    history.push("/createParty");
  }

  return (
    <Layout>
      <Header style={{ backgroundColor: 'whitesmoke' }}>ปาร์ตี้ทั้งหมด</Header>
      <Content>
        <Button type="primary" onClick={createParty}>สร้างปาร์ตี้</Button>
        <Card>
          {partyListData.map(data => (
            <PartyCard
              key={data.id}
              img={data.img}
              name={data.name}
              totalMember={data.totalMember}
              currentMember={data.currentMember}
              joinAction={() => { console.log('join ${data.name}') }}
            />
          ))}
        </Card>
      </Content>
    </Layout>
  )
}

export default PartyList
