package expenseservice

import (
	expensedomain "expense-management-system/cmd/expense/domain"
	expensedto "expense-management-system/cmd/expense/dto"
	expenseenum "expense-management-system/cmd/expense/enum"
	"expense-management-system/internal/contextkey"
	"expense-management-system/internal/mocks"
	"expense-management-system/pkg/jwt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func Test_expenseServiceImpl_formExpenseData(t *testing.T) {

	m := mocks.New()

	m.Config.ApprovalThreshold = 200000
	m.Config.MinExpenseAmount = 10000
	m.Config.MaxExpenseAmount = 50000000
	m.Config.SystemUserID = 1

	s := expenseServiceImpl{
		authrepository:    m.Authrepo,
		expenserepository: m.Expenserepo,
		transaction:       m.Transaction,
		validator:         m.Validator,
		cfg:               m.Config,
		jobClient:         m.JobClient,
	}

	app := fiber.New()
	fctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(fctx)

	fctx.Locals(contextkey.User, &jwt.AuthClaims{
		UserID: 10,
	})

	now := time.Now()

	type args struct {
		c   *fiber.Ctx
		req expensedto.CreateExpenseReq
		log *zap.Logger
	}
	tests := []struct {
		name    string
		s       *expenseServiceImpl
		args    args
		want    expensedomain.Expense
		wantErr bool
	}{
		{
			name: "Auto-approve flow",
			s:    &s,
			args: args{
				c:   fctx,
				log: zap.L(),
				req: expensedto.CreateExpenseReq{
					AmountIDR: 100000,
					Timestamp: now,
				},
			},
			want: expensedomain.Expense{
				UserID: 10,
				Amount: 100000,
				Status: expenseenum.ExpenseAutoApproved,
			},
			wantErr: false,
		},
		{
			name: "Awaiting approval when amount >= threshold",
			s:    &s,
			args: args{
				c: fctx,
				req: expensedto.CreateExpenseReq{
					AmountIDR: 200000,
					Timestamp: now,
				},
			},
			want: expensedomain.Expense{
				UserID: 10,
				Amount: 200000,
				Status: expenseenum.ExpenseAwaitingApproval,
			},
			wantErr: false,
		},
		{
			name: "Amount below minimum",
			s:    &s,
			args: args{
				c: fctx,
				req: expensedto.CreateExpenseReq{
					AmountIDR: 1,
					Timestamp: now,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := tt.s.formExpenseData(
				tt.args.c,
				tt.args.req,
				zap.L(),
			)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.UserID != tt.want.UserID {
				t.Errorf("UserID = %v, want %v", got.UserID, tt.want.UserID)
			}
			if got.Amount != tt.want.Amount {
				t.Errorf("Amount = %v, want %v", got.Amount, tt.want.Amount)
			}
			if got.Status != tt.want.Status {
				t.Errorf("Status = %v, want %v", got.Status, tt.want.Status)
			}
			if got.ID == uuid.Nil {
				t.Errorf("ID should be generated")
			}

			if got.Status == expenseenum.ExpenseAutoApproved {
				if got1 == nil {
					t.Errorf("Approval should not be nil")
				}
				if got2 == nil {
					t.Errorf("Payment should not be nil")
				}
				if got3 == nil {
					t.Errorf("Job task should not be nil")
				}
			}

			if got.Status == expenseenum.ExpenseAwaitingApproval {
				if got1 != nil || got2 != nil || got3 != nil {
					t.Errorf("No approval/payment/task should be created")
				}
			}
		})
	}
}
