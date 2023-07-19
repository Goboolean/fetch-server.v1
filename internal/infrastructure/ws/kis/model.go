package kis

import (
	"errors"
	"strings"

	"github.com/Goboolean/fetch-server/internal/infrastructure/ws"
)

type HeaderJson struct {
	ApprovalKey string `json:"approval_key"` // 실시간 접속키
	Custtype    string `json:"custtype"`     // 고객 타입 (P: 개인, B: 법인)
	TrType      string `json:"tr_type"`      // 거래 타입 (1. 등록, 2. 해제)
	ContentType string `json:"content_type"` // 컨텐츠 타입 (utf-8 고정)
}
type RequestBodyJson struct {
	Input RequestInputJson `json:"input"`
}

type RequestInputJson struct {
	TrId   string `json:"tr_id"`   // 거래 ID (H0STCNT0: 실시간 주식 체결가, H0STASP0: 주식 호가)
	TrCode string `json:"tr_code"` // 종목코드
}

type RequestJson struct {
	Header HeaderJson      `json:"header"`
	Body   RequestBodyJson `json:"body"`
}

type StockDetail struct {
	MKSC_SHRN_ISCD               string // 유가증권 단축 종목코드
	STCK_CNTG_HOUR               string // 주식 체결 시간
	STCK_PRPR                    string // 주식 현재가
	PRDY_VRSS_SIGN               string // 전일 대비 부호
	PRDY_VRSS                    string // 전일 대비
	PRDY_CTRT                    string // 전일 대비율
	WGHN_AVRG_STCK_PRC           string // 가중평균 주가
	STCK_OPRC                    string // 주식 시가
	STCK_HGPR                    string // 주식 고가
	STCK_LWPR                    string // 주식 저가
	ASKP1                        string // 매도호가1
	BIDP1                        string // 매수호가1
	CNTG_VOL                     string // 체결 거래량
	ACML_VOL                     string // 누적 거래량
	ACML_TR_PBMN                 string // 누적 거래대금
	SELN_CNTG_CSNU               string // 매도 체결 건수
	SHNU_CNTG_CSNU               string // 매수 체결 건수
	NTBY_CNTG_CSNU               string // 순매수 체결 건수
	CTTR                         string // 체결강도
	SELN_CNTG_SMTN               string // 총 매도 수량
	SHNU_CNTG_SMTN               string // 총 매수 수량
	CCLD_DVSN                    string // 체결구분
	SHNU_RATE                    string // 매수비율
	PRDY_VOL_VRSS_ACML_VOL_RATE  string // 전일 거래량 대비 등락률
	OPRC_HOUR                    string // 시가 시간
	OPRC_VRSS_PRPR_SIGN          string // 시가대비구분
	OPRC_VRSS_PRPR               string // 시가대비
	HGPR_HOUR                    string // 최고가 시간
	HGPR_VRSS_PRPR_SIGN          string // 고가대비구분
	HGPR_VRSS_PRPR               string // 고가대비
	LWPR_HOUR                    string // 최저가 시간
	LWPR_VRSS_PRPR_SIGN          string // 저가대비구분
	LWPR_VRSS_PRPR               string // 저가대비
	BSOP_DATE                    string // 영업 일자
	NEW_MKOP_CLS_CODE            string // 신 장운영 구분 코드
	TRHT_YN                      string // 거래정지 여부 (Y: 정지, N: 정상거래)
	ASKP_RSQN1                   string // 매도호가 잔량1
	BIDP_RSQN1                   string // 매수호가 잔량1
	TOTAL_ASKP_RSQN              string // 총 매도호가 잔량
	TOTAL_BIDP_RSQN              string // 총 매수호가 잔량
	VOL_TNRT                     string // 거래량 최전율
	PRDY_SMNS_HOUR_ACML_VOL      string // 전일 동시간 누적 거래량
	PRDY_SMNS_HOUR_ACML_VOL_RATE string // 전일 동시간 누적 거래량 비율
	HOUR_CLS_CODE                string // 시간 구분 코드
	MRKT_TRTM_CLS_CODE           string // 임의종료구분코드
	VI_STND_PRC                  string // 정적 VI 발동 기준가
}

func ToStockDetail(receivedString string) (*StockDetail, error) {
	data := strings.Split(receivedString, "^")
	if len(data) != 52 {
		return nil, errors.New("incorrect number of fields in response")
	}
	return &StockDetail{
		MKSC_SHRN_ISCD:               data[0],
		STCK_CNTG_HOUR:               data[1],
		STCK_PRPR:                    data[2],
		PRDY_VRSS_SIGN:               data[3],
		PRDY_VRSS:                    data[4],
		PRDY_CTRT:                    data[5],
		WGHN_AVRG_STCK_PRC:           data[6],
		STCK_OPRC:                    data[7],
		STCK_HGPR:                    data[8],
		STCK_LWPR:                    data[9],
		ASKP1:                        data[10],
		BIDP1:                        data[11],
		CNTG_VOL:                     data[12],
		ACML_VOL:                     data[13],
		ACML_TR_PBMN:                 data[14],
		SELN_CNTG_CSNU:               data[15],
		SHNU_CNTG_CSNU:               data[16],
		NTBY_CNTG_CSNU:               data[17],
		CTTR:                         data[18],
		SELN_CNTG_SMTN:               data[19],
		SHNU_CNTG_SMTN:               data[20],
		CCLD_DVSN:                    data[21],
		SHNU_RATE:                    data[22],
		PRDY_VOL_VRSS_ACML_VOL_RATE:  data[23],
		OPRC_HOUR:                    data[24],
		OPRC_VRSS_PRPR_SIGN:          data[25],
		OPRC_VRSS_PRPR:               data[26],
		HGPR_HOUR:                    data[27],
		HGPR_VRSS_PRPR_SIGN:          data[28],
		HGPR_VRSS_PRPR:               data[29],
		LWPR_HOUR:                    data[30],
		LWPR_VRSS_PRPR_SIGN:          data[31],
		LWPR_VRSS_PRPR:               data[32],
		BSOP_DATE:                    data[33],
		NEW_MKOP_CLS_CODE:            data[34],
		TRHT_YN:                      data[35],
		ASKP_RSQN1:                   data[36],
		BIDP_RSQN1:                   data[37],
		TOTAL_ASKP_RSQN:              data[38],
		TOTAL_BIDP_RSQN:              data[39],
		VOL_TNRT:                     data[40],
		PRDY_SMNS_HOUR_ACML_VOL:      data[41],
		PRDY_SMNS_HOUR_ACML_VOL_RATE: data[42],
		HOUR_CLS_CODE:                data[43],
		MRKT_TRTM_CLS_CODE:           data[44],
		VI_STND_PRC:                  data[45],
	}, nil
}

func ToStockAggsDetail(s *StockDetail) (*ws.StockAggregate, error) {
	return &ws.StockAggregate{}, nil
}
