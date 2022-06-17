import { Layout, Button, Row, Col } from 'antd';
import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import ProductCard from '../components/productCard';
import { route } from '../config';
import * as productService from '../services/product'

const { Header, Content } = Layout;

function ProductList() {
  let history = useHistory();
  const [productListData, setProductListData] = useState<any[]>();

  useEffect(() => {
    // getProducts();
  }, []);

  const getProducts = async () => {
    const list = await productService.getList();
    setProductListData(list);
  };

  return (
      <Content style={{ backgroundColor: 'whitesmoke' }}>
        <Row gutter={[16, 16]} style={{ margin: '10px' }}>
          {productListData
            ? productListData.map((data) => (
                <Col span={12} key={data.id}>
                  <ProductCard
                    key={data.ID}
                    img={""}
                    name={data.Name}
                    price={data.Price}
                    action={() => {}}
                  />
                </Col>
              ))
            : null}
        </Row>
      </Content>
  );
}

export default ProductList;
